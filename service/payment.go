package service

import (
	"context"
	"database/sql"
	"fmt"
	pb "payment_service/genproto/payments"
	pbu "payment_service/genproto/reservations"
	"payment_service/storage/postgres"
	"testing"
)

type PaymentService struct {
	pb.UnimplementedPaymentsServer
	Payment           *postgres.PaymentRepo
	ReservationCleint pbu.ReservationServiceClient
}

func NewPaymentService(db *sql.DB) *PaymentService {
	t := testing.T{}
	return &PaymentService{
		Payment: postgres.NewPaymentRepo(db, &t),
	}
}

func (p *PaymentService) GetPayments(ctx context.Context, req *pb.PaymentsFilter) (*pb.AllPayments, error) {
	r, err := p.Payment.GetPayments(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (p *PaymentService) UpdatePayment(ctx context.Context, req *pb.Payment) (*pb.Void, error) {

	exist, err := p.ReservationCleint.ValidateReservationId(ctx, &pbu.Id{Id: req.ReservationId})
	if !exist.Exists || err != nil {
		return nil, fmt.Errorf("reservation id invalid")
	}

	_, err = p.Payment.UpdatePayment(req)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentService) DeletePayment(ctx context.Context, rep *pb.Id) (*pb.Void, error) {
	_, err := p.Payment.DeletePayment(rep)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentService) GetStatus(ctx context.Context, req *pb.Id) (*pb.Status, error) {
	r, err := p.Payment.GetStatus(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (p *PaymentService) MakePayment(ctx context.Context, req *pb.Payment) (*pb.Id, error) {
	exist, err := p.ReservationCleint.ValidateReservationId(ctx, &pbu.Id{Id: req.ReservationId})
	if !exist.Exists || err != nil {
		return nil, fmt.Errorf("reservation id invalid")
	}

	r, err := p.Payment.MakePayment(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// func (p *PaymentService) ValidatePaymentId(ctx context.Context, req *pb.Id) (*pb.Exists, error) {
// 	exist, err := p.Payment.ValidatePaymentId(req)
// 	if !exist.Exists || err != nil {
// 		return &pb.Exists{Exists: exist.Exists}, err
// 	}
// 	return &pb.Exists{Exists: exist.Exists}, nil
// }
