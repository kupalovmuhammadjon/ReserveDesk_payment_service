package postgres

import (
	"database/sql"
	pb "payment_service/genproto/payments"
	"testing"
)

func Test_MakePayment(t *testing.T) {
	db := sql.DB{}
	payments := NewPaymentRepo(&db, t)

	payment := pb.Payment{
		Id:            "1",
		ReservationId: "2",
		PaymentMethod: "adsd",
		Amount:        112321,
	}

	_, err := payments.MakePayment(&payment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new payment", err)
	}
}

func Test_UpdatePayment(t *testing.T) {
	db := sql.DB{}
	payments := NewPaymentRepo(&db, t)

	payment := pb.Payment{
		Id:            "1",
		ReservationId: "2",
		PaymentMethod: "adsd",
		Amount:        112321,
	}
	_, err := payments.UpdatePayment(&payment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when updating a payment", err)
	}
}

func Test_GetPayment(t *testing.T) {
	db := sql.DB{}
	payments := NewPaymentRepo(&db, t)

	payment := pb.PaymentsFilter{
		PaymentsTo:   1223,
		PaymentsFrom: 123223,
	}
	_, err := payments.GetPayments(&payment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when getting a payment", err)
	}
}

func Test_DeletePayment(t *testing.T) {
	db := sql.DB{}
	payments := NewPaymentRepo(&db, t)

	payment := pb.Id{
		Id: "1",
	}

	_, err := payments.DeletePayment(&payment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when updating a payment", err)
	}
}

func TestPaymentRepo_GetStatus(t *testing.T) {
	db := sql.DB{}
	payments := NewPaymentRepo(&db, t)

	payment := pb.Id{
		Id: "1",
	}

	_, err := payments.GetStatus(&payment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when getting a payment", err)
	}
}

func Test_ValidatePaymentId(t *testing.T) {
	db := sql.DB{}
	payments := NewPaymentRepo(&db, t)

	payment := pb.Id{
		Id: "1",
	}

	_, err := payments.ValidatePaymentId(&payment)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when validating a payment", err)
	}
}
