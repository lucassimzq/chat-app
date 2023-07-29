package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "log"

    "github.com/joho/godotenv"
    
    "github.com/zhenquansim/chat-app/pkg/websocket"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadGoDotEnvVariable() {

    // load .env file
    err := godotenv.Load(".env")
  
    if err != nil {
      log.Fatalf("Error loading .env file")
    }
}

func serveWs(db *mongo.Database, pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
        Db:   db,
    }

    pool.Register <- client
    client.Read()
}

func setupRoutes(db *mongo.Database) {
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(db, pool, w, r)
    })
}

func main() {
    loadGoDotEnvVariable()

    var uri = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", 
        os.Getenv("DATABASE_USERNAME"),
        os.Getenv("DATABASE_PASSWORD"),
        os.Getenv("DATABASE_HOST"))

    // Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    client, err := mongo.Connect(context.TODO(), opts)

    if err != nil {
        panic(err)
    }
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

    // Send a ping to confirm a successful connection
    var result bson.M
    if err := client.Database("main").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
        panic(err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

    fmt.Println("Distributed Chat App v0.01")
    setupRoutes(client.Database("chat-app"))
    http.ListenAndServe(":8080", nil)
}