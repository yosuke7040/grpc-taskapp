package router

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rs/cors"
	"github.com/yosuke7040/grpc-taskapp/backend/handler"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/interceptor"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/model/db"
	sqlcRepo "github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/sqlc/repository"
	"github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/auth/v1/auth_v1connect"
	"github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/task/v1/task_v1connect"
	"github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/user/v1/user_v1connect"
	"github.com/yosuke7040/grpc-taskapp/backend/usecase"
	"github.com/yosuke7040/grpc-taskapp/backend/utils/auth"
	"github.com/yosuke7040/grpc-taskapp/backend/utils/contextkey"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type serverMuxEngine struct {
	router     *http.ServeMux
	port       Port
	ctxTimeout time.Duration
}

func newServerMuxEngine(
	port Port,
	t time.Duration,
) *serverMuxEngine {
	return &serverMuxEngine{
		router:     http.NewServeMux(),
		port:       port,
		ctxTimeout: t,
	}
}

func (s *serverMuxEngine) Listen() {
	slog.Info("Server is running", "port", s.port)

	dbConn, err := sql.Open("mysql", "gogo:gogo@tcp(mysql:3306)/app_db?parseTime=true")
	if err != nil {
		slog.Error("failed to open database", "err", err)
	}

	err = dbConn.Ping()
	if err != nil {
		slog.Error("failed to ping database", "err", err)
	}

	defer dbConn.Close()
	qry := db.New(dbConn)

	s.setupHandlers(qry)

	handler := cors.AllowAll().Handler(h2c.NewHandler(s.router, &http2.Server{}))

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.port),
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown: ", err)
		os.Exit(1)
	}

	slog.Info("Server shutdown properly")
}

func (s *serverMuxEngine) setupHandlers(qry db.Querier) {
	authInterceptor := connect.WithInterceptors(interceptor.NewAuthInterceptor("github.com/yosuke7040/grpc-taskapp/backend", "/home/vscode/.grpc-taskapp/id_rsa"))

	s.router.Handle(auth_v1connect.NewAuthServiceHandler(s.buildAuthServerHandler(qry)))
	s.router.Handle(user_v1connect.NewUserServiceHandler(s.buildUserServerHandler(qry)))
	s.router.Handle(task_v1connect.NewTaskServiceHandler(s.buildTaskServerHandler(qry), authInterceptor))
}

func (s *serverMuxEngine) buildAuthServerHandler(qry db.Querier) *handler.AuthHandler {
	tm, _ := auth.NewTokenManager("github.com/yosuke7040/grpc-taskapp/backend", "/home/vscode/.grpc-taskapp/id_rsa")
	// if err != nil {
	// 	return nil, err
	// }

	repo := sqlcRepo.NewUserRepository(qry)
	uc := usecase.NewAuthUsecase(repo, tm, (1 * time.Hour))
	server := handler.NewAuthHandler(uc)

	return &server
}

func (s *serverMuxEngine) buildUserServerHandler(qry db.Querier) *handler.UserHandler {
	repo := sqlcRepo.NewUserRepository(qry)
	uc := usecase.NewUserUsecase(repo)
	server := handler.NewUserHandler(uc)

	return &server
}

func (s *serverMuxEngine) buildTaskServerHandler(qry db.Querier) *handler.TaskHandler {
	cr := contextkey.NewContextReader()
	repo := sqlcRepo.NewTaskRepository(qry)
	uc := usecase.NewTaskUsecase(repo)
	server := handler.NewTaskHandler(uc, cr)

	return &server
}
