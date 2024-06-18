package main

// "context"
// "database/sql"
// "log/slog"
// "net/http"

// "github.com/yosuke7040/grpc-taskapp/infrastructure"

// _ "github.com/go-sql-driver/mysql"
// "github.com/rs/cors"
// "github.com/yosuke7040/grpc-taskapp/infrastructure/persistence/model/db"
// "golang.org/x/net/http2"
// "golang.org/x/net/http2/h2c"

import (
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure"
)

func main() {
	app := infrastructure.NewConfig()

	app.WebServerPort("8080").
		WebServer().
		Start()

	// if err := run(); err != nil {
	// 	slog.Error("Error: %v", err)
	// }

	// func run() error {
	// 	ctx, cancel := context.WithCancel(context.Background())
	// 	defer cancel()

	// 	cfg, err := sql.Open("mysql", "user:password@/dbname?parseTime=ture")
	// 	if err != nil {
	// 		return slog.Error("failed to open database: %v", err)
	// 	}

	// 	qry := db.New(cfg)

	// 	taskServer := di.InitTask(qry)

	// 	mux := http.NewServeMux()
	// 	mux.Handle(task_v1connect.NewTaskServiceHandler(taskServer))

	// return http.ListenAndServe(
	//
	//	"localhost:8080",
	//	cors.AllowAll().Handler(
	//		h2c.NewHandler(mux, &http2.Server{}),
	//	),
	//
	// )
}
