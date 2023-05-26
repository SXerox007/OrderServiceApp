package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	so "github.com/SXerox007/OrderServiceApp/api/services/order"
	sp "github.com/SXerox007/OrderServiceApp/api/services/product"
	"github.com/SXerox007/OrderServiceApp/base/server"
	"github.com/SXerox007/OrderServiceApp/protos/order"
	"github.com/SXerox007/OrderServiceApp/protos/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ServerSetup()
}

//brain setup
func ServerSetup() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Error in server start:", err)
		return
	}
	//Create the New gRPC Server
	srv := server.CreateNewgRPCServer(false, nil)
	//Register reflection on gRPC server
	reflection.Register(srv)
	// all the rpc services
	setupServices(srv)

	go func() {
		fmt.Println("Server start on Port:50051")
		if err := srv.Serve(listner); err != nil {
			log.Println("Error in Serve:", err)
			return
		}
	}()
	//make a channnel that will wait for server to close or interrupt
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// this will block the code while server
	<-ch
	log.Println("Exit the Server with:", os.Interrupt)
}

// All the services
func setupServices(srv *grpc.Server) {

	psc := sp.New("ADMIN_JOSH", make(map[string]*product.Product))
	psc.RegisterProductService(srv)

	osc := so.New("ADMIN_JOSH", make(map[string]*order.Order), psc)
	osc.RegisterOrderService(srv)
}
