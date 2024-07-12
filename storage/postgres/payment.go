package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	pb "payment_service/genproto/payments"
	"testing"
	"time"
)

type PaymentRepo struct {
	Db *sql.DB
	pb.UnimplementedPaymentsServer
	t *testing.T
}

func NewPaymentRepo(db *sql.DB, t *testing.T) *PaymentRepo {
	return &PaymentRepo{
		Db: db,
		t:  t,
	}
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

func (p *PaymentRepo) UpdatePayment(req *pb.Payment) (*pb.Void, error) {
	_, err := p.Db.Exec("update payments set reservation_id=$1, amount=$2, payment_method=$3 updated_at=$4 where id=$5", req.ReservationId, req.PaymentMethod, req.Amount, time.Now(), req.Id)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentRepo) DeletePayment(rep *pb.Id) (*pb.Void, error) {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.Void{}, err
	}
	_, err = p.Db.Exec("update payments set delete_at=$1 where id=$2", time.Now(), rep.Id)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentRepo) GetPayments(req *pb.PaymentsFilter) (*pb.AllPayments, error) {
	var params []interface{}

	query := "select id, reservation_id, amount, payment_method, payment_status, created_at  from payments where deleted_at is null"

	if req.PaymentsFrom > 0 {
		params = append(params, req.PaymentsFrom)
		query = fmt.Sprintf("payment_method = $%d", len(params))
	}

	if req.PaymentsTo > 0 {
		params = append(params, req.PaymentsTo)
		query = fmt.Sprintf("payment_status = $%d", len(params))
	}

	if req.Limit > 0 {
		params = append(params, req.Limit)
		query = fmt.Sprintf(" LIMIT = $%d", len(params))
	}

	if req.Offset > 0 {
		params = append(params, req.Offset)
		query = fmt.Sprintf(" OFFSET = $%d", len(params))
	}

	rows, err := p.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	var payments pb.AllPayments
	for rows.Next() {
		var payment pb.AllPayment
		err := rows.Scan(&payment.Id, &payment.ReservationId, &payment.Amount, &payment.PaymentMethod, &payment.CreatedAt)
		if err != nil {
			return nil, err
		}
		payments.AllPayments = append(payments.AllPayments, &payment)
	}
	return &payments, nil
}

func (p *PaymentRepo) GetStatus(req *pb.Id) (*pb.Status, error) {
	_, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return nil, err
	}
	resp := pb.Status{}
	err = p.Db.QueryRow("select payment_status from payments where deleted_at is null and id=$1", req.Id).Scan(&resp.PaymentStatus)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
