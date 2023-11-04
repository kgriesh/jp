package main

import (
	"jp/app"
	"jp/app/db"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":8000"
	//listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	log.Fatalf("Error occurred: %s", err.Error())
	// }
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	database, err := db.NewDbService(dbUser, dbPassword, dbName)
	if err != nil {
		log.Printf("error %v", err)
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.GetConnection().Close()

	handler := app.New(database)
	http.ListenAndServe(addr, handler)
	// server := &http.Server{
	// 	Handler: handler,
	// }
	// go func() {
	// 	server.Serve(listener)
	// }()
	// defer Stop(server)
	// log.Printf("Started server on %s", addr)

	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// log.Println(fmt.Sprint(<-ch))
	// log.Println("Stopping API server.")
}

// func Stop(server *http.Server) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := server.Shutdown(ctx); err != nil {
// 		log.Printf("Server shut down error: %v\n", err)
// 		os.Exit(1)
// 	}
// }
