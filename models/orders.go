package models

import (
	"errors"
	// "fmt"
	"github.com/jackyyf/gook/db"
	"github.com/jackyyf/gook/utils/log"
)

type OrderStatus int

const (
	STAT_NEW    OrderStatus = iota
	STAT_PAID   OrderStatus = iota
	STAT_DONE   OrderStatus = iota
	STAT_CANCEL OrderStatus = -1
)

func (s OrderStatus) String() string {
	if s == STAT_NEW {
		return "New"
	} else if s == STAT_PAID {
		return "Paid"
	} else if s == STAT_DONE {
		return "Done"
	} else {
		return "UNKNOWN_STATUS"
	}
}

type OrderIn struct {
	id      int32
	Book    *Book
	Amount  int32
	Price   float64
	Status  OrderStatus
	BillRef *Bill
}

type OrderOut struct {
	id      int32
	Book    *Book
	Amount  int32
	Price   float64
	BillRef *Bill
}

func (order *OrderIn) ID() int32 {
	return order.id
}

func (order *OrderOut) ID() int32 {
	return order.id
}

func (order *OrderIn) Create() (err error) {
	if order == nil {
		err = errors.New("Error: unable to create nil order in.")
		log.Alert(err.Error())
		return
	}
	if order.Book == nil {
		err = errors.New("Error: no book ref set.")
		log.Alert(err.Error())
		return
	}
	err = db.QueryRow(
		`INSERT INTO "orderin" (book, amount, price, status, billref) VALUES ($1, $2, $3, 0, NULL) RETURNING id`,
		order.Book.ID(), order.Amount, order.Price).Scan(&order.id)
	if err != nil {
		log.Error("Create order in failed: %s", err)
		return
	}
	log.Info("New order in(ID=%d) created successfully.", order.id)
	return
}

func (order *OrderIn) Cancel() (err error) {
	if order == nil {
		err = errors.New("Error: unable to cancel nil order in.")
		log.Alert(err.Error())
		return
	}
	if order.Status != STAT_CANCEL {
		err = errors.New("Error: can't cancel paid order.")
		log.Alert(err.Error())
		return
	}
	res, err := db.Exec(`UPDATE "orderin" SET status=-1 WHERE id=$1`, order.ID())
	if err != nil {
		log.Error("Update order in failed: %s", err)
		return
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting changes to an unexist order!")
		log.Alert(err.Error())
		return
	} else {
		log.Info("Order in(ID=%d) cancelled.", order.ID())
		return
	}
}

func (order *OrderIn) Pay(bill *Bill) (err error) {
	if order == nil {
		err = errors.New("Error: unable to create nil order in.")
		log.Alert(err.Error())
		return
	}
	if order.Status != STAT_NEW {
		err = errors.New("Error: order already paid.")
		log.Alert(err.Error())
		return
	}
	if bill == nil {
		err = errors.New("Error: nil bill")
		log.Alert(err.Error())
		return
	}
	tx, err := db.Transaction()
	if err != nil {
		log.Crit("Unable to create transaction: %s", err)
		return
	}
	err = bill.Create(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	res, err := tx.Exec(
		`UPDATE "orderin" SET status=1, billref=$1`,
		order.BillRef.ID())
	if err != nil {
		log.Error("Update order in failed: %s", err)
		tx.Rollback()
		return
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting changes to an unexist order!")
		log.Alert(err.Error())
		tx.Rollback()
		return
	} else {
		log.Info("Order in(ID=%d) paid. Bill=%d", order.ID(), order.BillRef.ID())
		tx.Commit()
		return
	}
}

func (order *OrderIn) Finish() (err error) {
	if order == nil {
		err = errors.New("Error: unable to create nil order in.")
		log.Alert(err.Error())
		return
	}
	if order.Status != STAT_PAID {
		err = errors.New("Error: order not ready to finalize.")
		log.Alert(err.Error())
		return
	}
	tx, err := db.Transaction()
	if err != nil {
		log.Crit("Unable to create transaction: %s", err)
		return
	}
	res, err := tx.Exec(
		`UPDATE "orderin" SET status=2`,
		order.BillRef.ID())
	if err != nil {
		log.Error("Update order in failed: %s", err)
		tx.Rollback()
		return
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting changes to an unexist order!")
		tx.Rollback()
		log.Alert(err.Error())
		return
	} else {
		log.Info("Order in(ID=%d) paid. Bill=%d", order.ID(), order.BillRef.ID())
	}
	res, err = tx.Exec(
		`UPDATE "book" SET amount=amount+$1 WHERE id=$2`, order.Amount, order.Book.ID())
	if err != nil {
		log.Error("Update book failed: %s", err)
		tx.Rollback()
		return
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting changes to an unexist book!")
		tx.Rollback()
		log.Alert(err.Error())
		return
	} else {
		log.Info("Book amount updated.", order.ID(), order.BillRef.ID())
	}
	tx.Commit()
	return
}

func (order *OrderOut) Create() (err error) {
	if order == nil {
		err = errors.New("Error: unable to create nil order out.")
		log.Alert(err.Error())
		return
	}
	if order.Book == nil {
		err = errors.New("Error: no book ref set.")
		log.Alert(err.Error())
		return
	}
	if order.BillRef == nil {
		err = errors.New("Error: no bill ref set.")
		log.Alert(err.Error())
		return
	}
	tx, err := db.Transaction()
	if err != nil {
		log.Crit("Unable to create transaction: %s", err)
		return
	}
	err = tx.QueryRow(
		`INSERT INTO "orderout" (book, amount, price, billref) VALUES ($1, $2, $3, $4) RETURNING id`,
		order.Book.ID(), order.Amount, order.Price, order.BillRef.ID()).Scan(&order.id)
	if err != nil {
		log.Error("Create order out failed: %s", err)
		tx.Rollback()
		return
	}
	log.Info("New order out(ID=%d) created successfully.", order.id)
	res, err := tx.Exec(`UPDATE "book" SET amount=amount-$1 WHERE id=$2`, order.Amount, order.Book.ID())
	if err != nil {
		log.Error("Update book failed: %s", err)
		tx.Rollback()
		return
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting changes to an unexist book!")
		tx.Rollback()
		log.Alert(err.Error())
		return
	} else {
		log.Info("Book amount updated.", order.ID(), order.BillRef.ID())
		tx.Commit()
	}
	return
}
