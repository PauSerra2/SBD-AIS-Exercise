package server

import (
    "context"
    "exc8/pb/pb"
    "fmt"
    "log/slog"
    "net"

    "google.golang.org/grpc"
    emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// GRPCService implements the OrderServiceServer interface
type GRPCService struct {
    pb.UnimplementedOrderServiceServer
    drinks []*pb.Drink
    orders map[int32]int32 // drink_id -> total quantity
}

// StartGrpcServer starts the gRPC server on port 4000
func StartGrpcServer() error {
    srv := grpc.NewServer()
    grpcService := NewGRPCService()
    pb.RegisterOrderServiceServer(srv, grpcService)

    lis, err := net.Listen("tcp", ":4000")
    if err != nil {
        return err
    }
    slog.Info("gRPC server listening on :4000")

    if err := srv.Serve(lis); err != nil {
        return err
    }
    return nil
}

// todo implement functions

func NewGRPCService() *GRPCService {
    return &GRPCService{
        drinks: []*pb.Drink{
            {Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
            {Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
            {Id: 3, Name: "Coffee", Price: 0, Description: "Mifare isn't that secure"},
        },
        orders: make(map[int32]int32),
    }
}

// GetDrinks returns all available drinks
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.DrinksResponse, error) {
    return &pb.DrinksResponse{Drinks: s.drinks}, nil
}

// OrderDrink stores an order in memory and returns a success flag
func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    order := req.Order
    s.orders[order.DrinkId] += order.Quantity
    fmt.Sprintf("Ordered %d of drink_id %d", order.Quantity, order.DrinkId)
    return &pb.OrderResponse{Success: true}, nil
}

// GetOrders returns all accumulated orders
func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrdersResponse, error) {
    var orders []*pb.Order
    for id, qty := range s.orders {
        orders = append(orders, &pb.Order{DrinkId: id, Quantity: qty})
    }
    return &pb.OrdersResponse{Orders: orders}, nil
}
