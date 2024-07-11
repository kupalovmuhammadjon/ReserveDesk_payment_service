package main

import (
	"fmt"
	"log"
	"net"
	"payment_service/config"

	pb "payment_service/genproto/payments"
	"payment_service/pkg/logger"
	"payment_service/service"
	"payment_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	logger, err := logger.New("debug", "develop", "app.log")
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := postgres.ConnectDB()
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	fmt.Println(logger)
	fmt.Println("wsdfghn")

	listener, err := net.Listen("tcp", cfg.PAYMENT_SERVICE_PORT)
	if err != nil {
		logger.Fatal("error while making tcp connection")
		return
	}

	server := grpc.NewServer()

	pb.RegisterPaymentsServer(server, service.NewPaymentService(db))

	err = server.Serve(listener)
	if err != nil {
		return
	}
}
