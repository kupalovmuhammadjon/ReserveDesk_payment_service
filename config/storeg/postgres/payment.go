package postgres

import (
	"database/sql"
	pb "payment_service/genproto/payments"
	"time"
)


type PaymentRepo struct {
	Db *sql.DB
	pb.UnimplementedPaymentsServer
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{Db: db}
}

func (p *PaymentRepo) MakePayment(req *pb.Payment) (*pb.Id, error) {
	_, err := p.Db.Exec("insert into payments(reservation_id, amount, payment_metod, created_at) values($1, $2, $3, $4)", req.ReservationId, req.Amount, req.PaymentMethod, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Id{Id: req.ReservationId}, nil
}

func (p *PaymentRepo) ValidatePaymentId(id *pb.Id) (*pb.Exists, error) {
	query := `
		select 
			case 
				whene id = $1 then true
			else 
				false
			end
		from 
			payments
		where id = $1 deleted_at is null 
		`
		res := pb.Exists{}
		err := p.Db.QueryRow(query, id).Scan(&res.Exists)
		return &res, err
}