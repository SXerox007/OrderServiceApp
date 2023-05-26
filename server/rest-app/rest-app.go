package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/SXerox007/OrderServiceApp/protos/order"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	authpoint = flag.String("auth_end_points", "localhost:50051", "expose josh end point ")
)

func ExposePoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	// RegisterGateway grpcgw
	err := order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)

	if err != nil {
		return err
	}
	grpcMux := http.NewServeMux()
	grpcMux.Handle("/", mux)
	//grpcMux.HandleFunc("/swagger/", serveSwagger)
	log.Println("Starting Endpoint Exposed Server: localhost:5051")
	http.ListenAndServe(address, grpcMux)
	return nil
}

func main() {
	if err := ExposePoint(":5051"); err != nil {
		log.Fatal("Error in serve", err)
	}
}
