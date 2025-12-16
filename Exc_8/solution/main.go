package main

import (
    "log"
    "time"

    "exc8/server"
    "exc8/client"
)

func main() {
    // Start the gRPC server 
    go func() {
        if err := server.StartGrpcServer(); err != nil {
            log.Fatalf("failed to start gRPC server: %v", err)
        }
    }()

    // Give the server a moment to start
    time.Sleep(1 * time.Second)

    // Create and run the gRPC client
    c, err := client.NewGrpcClient()
    if err != nil {
        log.Fatalf("failed to create gRPC client: %v", err)
    }
    if err := c.Run(); err != nil {
        log.Fatalf("client run failed: %v", err)
    }

}
