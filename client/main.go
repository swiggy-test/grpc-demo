package main

import (
    "google.golang.org/grpc"
    "log"
    pb "github.com/swiggy-test/grpc-demo/models"
    "context"
    "time"
)

func main() {
    conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Error connecting to server: %v", err.Error())
    }
    defer conn.Close()

    client := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    SayHello(client, ctx)
    GreetFullName(client, ctx)
}

func GreetFullName(client pb.GreeterClient, ctx context.Context) {
    r, err := client.GreetFullName(ctx, &pb.Name{FirstName: "Harry", LastName: "Potter"})
    if err != nil {
        log.Fatalf("Can't spell your name :'( : %v", err)
    }
    log.Printf("%v", r)
}

func SayHello(client pb.GreeterClient, ctx context.Context) {
    r, err := client.SayHello(ctx, &pb.Name{FirstName: "Harry", LastName: "Potter"})
    if err != nil {
        log.Fatalf("Can't get a hello :'( : %v", err)
    }
    log.Printf("Greeting: %s", r.GetResponse())
}