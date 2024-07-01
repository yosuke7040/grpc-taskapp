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

	_ "github.com/go-sql-driver/mysql"

	"github.com/rs/cors"
	userUsecase "github.com/yosuke7040/grpc-taskapp/backend/app/user"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/model/db"
	sqlcRepo "github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/sqlc/repository"
	"github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/user/v1/user_v1connect"
	userHandler "github.com/yosuke7040/grpc-taskapp/backend/presentation/user"
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
	// s.router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	// path, handler := s.buildUserServerHandler(qry)
	// s.router.Handle(path, handler)

	// s.router.Handle(user_v1connect.NewUserServiceHandler(s.buildUserServerHandler(qry)))

	path, userHandler := user_v1connect.NewUserServiceHandler(s.buildUserServerHandler(qry))
	s.router.Handle(path, userHandler)
}

func (s *serverMuxEngine) buildUserServerHandler(qry db.Querier) *userHandler.Handler {
	// func (s *serverMuxEngine) buildUserServerHandler(qry db.Querier) (string, http.Handler) {
	repo := sqlcRepo.NewUserRepository(qry)
	// srv := service.NewUserService(repo)
	uc := userUsecase.NewFindUserUseCase(repo)
	server := userHandler.NewHandler(uc)

	return &server
	// return user_v1connect.NewUserServiceHandler(&server)
}
