package client

import (
    "context"
    "exc8/pb/pb"
    "fmt"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
    client pb.OrderServiceClient
}

// NewGrpcClient creates a new gRPC client connected to localhost:4000
func NewGrpcClient() (*GrpcClient, error) {
    conn, err := grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }
    client := pb.NewOrderServiceClient(conn)
    return &GrpcClient{client: client}, nil
}

// Run executes the client workflow: list drinks, order drinks, order more, and get totals
func (c *GrpcClient) Run() error {
    ctx := context.Background()
    // todo
    // 1. List drinks
    fmt.Println("Requesting drinks 🍹🍺☕")
    fmt.Println("Available drinks:")
    drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
    if err != nil {
        return err
    }
    for _, d := range drinksResp.Drinks {
        fmt.Printf("\t> id:%d  name:\"%s\"  price:%d  description:\"%s\"\n",
            d.Id, d.Name, d.Price, d.Description)
    }

    // 2. Order a few drinks
    fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
    fmt.Println("\t> Ordering: 2 x Spritzer")
    _, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{Order: &pb.Order{DrinkId: 1, Quantity: 2}})
    fmt.Println("\t> Ordering: 2 x Beer")
    _, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{Order: &pb.Order{DrinkId: 2, Quantity: 2}})
    fmt.Println("\t> Ordering: 2 x Coffee")
    _, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{Order: &pb.Order{DrinkId: 3, Quantity: 2}})

    // 3. Order more drinks
    fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
    fmt.Println("\t> Ordering: 6 x Spritzer")
    _, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{Order: &pb.Order{DrinkId: 1, Quantity: 6}})
    fmt.Println("\t> Ordering: 6 x Beer")
    _, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{Order: &pb.Order{DrinkId: 2, Quantity: 6}})
    fmt.Println("\t> Ordering: 6 x Coffee")
    _, _ = c.client.OrderDrink(ctx, &pb.OrderRequest{Order: &pb.Order{DrinkId: 3, Quantity: 6}})

    // 4. Get order total
    fmt.Println("Getting the bill 💹💹💹")
    ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
    if err != nil {
        return err
    }
    for _, o := range ordersResp.Orders {
        var name string
        for _, d := range drinksResp.Drinks {
            if d.Id == o.DrinkId {
                name = d.Name
                break
            }
        }
        fmt.Printf("\t> Total: %d x %s\n", o.Quantity, name)
    }

    fmt.Println("Orders complete!")
    return nil
}
