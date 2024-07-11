package service

import (
	"context"
	"database/sql"
	pb "payment_service/genproto/payments"
	"payment_service/storage/postgres"
)

type PaymentService struct {
	pb.UnimplementedPaymentsServer
	Repo *postgres.PaymentRepo
}

func NewPaymentService(db *sql.DB) *PaymentService {
	return &PaymentService{
		Repo: postgres.NewPaymentRepo(db),
	}
}

func (p *PaymentService) GetPayments(ctx context.Context, req *pb.PaymentsFilter) (*pb.AllPayments, error) {
	r, err := p.Repo.GetPayments(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (p *PaymentService) UpdatePayment(ctx context.Context, req *pb.Payment) (*pb.Void, error) {
	_, err := p.Repo.UpdatePayment(req)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentService) DeletePayment(ctx context.Context, rep *pb.Id) (*pb.Void, error) {
	_, err := p.Repo.DeletePayment(rep)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentService) GetStatus(ctx context.Context, req *pb.Id) (*pb.Status, error) {
	r, err := p.Repo.GetStatus(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}


func (p *PaymentService) MakePayment(ctx context.Context, rep *pb.Payment) (*pb.Id, error) {
	exist1, err := p.Repo.V
	r, err := p.MakePayment(rep)
	if err != nil {
		return nil, err
	}
	return r, nil
}
