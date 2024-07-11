package pkg

import (
	"errors"
	"log"
	"payment_service/config"
	 pbu "payment_service/genproto/payments"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateReservationClient(cfg *config.Config) pbu.PaymentsClient {
	conn, err := grpc.NewClient(cfg.PAYMENT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbu.NewPaymentsClient(conn)
}
