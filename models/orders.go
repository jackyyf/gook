package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackyyf/gook/db"
	"github.com/jackyyf/gook/utils/log"
	"time"
)

type OrderStatus int

const (
	STAT_NEW    OrderStatus = 0
	STAT_PAID   OrderStatus = 1
	STAT_DONE   OrderStatus = 2
	STAT_CANCEL OrderStatus = -1
)

type _status struct {
	STAT_NEW    OrderStatus
	STAT_PAID   OrderStatus
	STAT_DONE   OrderStatus
	STAT_CANCEL OrderStatus
}

var Status = _status{STAT_NEW: STAT_NEW, STAT_PAID: STAT_PAID, STAT_DONE: STAT_DONE, STAT_CANCEL: STAT_CANCEL}

func (s OrderStatus) String() string {
	if s == STAT_NEW {
		return "New"
	} else if s == STAT_PAID {
		return "Paid"
	} else if s == STAT_DONE {
		return "Done"
	} else if s == STAT_CANCEL {
		return "Cancelled"
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
	tx, err := db.Transaction()
	if err != nil {
		log.Crit("Unable to create transaction: %s", err)
		return
	}
	if bill == nil {
		order.BillRef = new(Bill)
		order.BillRef.Amount = float64(order.Amount) * order.Price
		err = order.BillRef.Create(tx)
		if err != nil {
			log.Error("Create bill failed: %s", err)
			tx.Rollback()
			return
		}
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
	tx, err := db.Transaction()
	if err != nil {
		log.Crit("Unable to create transaction: %s", err)
		return
	}
	if order.BillRef == nil {
		order.BillRef = new(Bill)
		order.BillRef.Amount = float64(order.Amount) * order.Price
		err = order.BillRef.Create(tx)
		if err != nil {
			log.Error("Create bill failed: %s", err)
			tx.Rollback()
			return
		}
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

func GetOrderIns(book *Book) (ret []OrderIn, err error) {
	query := `SELECT orderin.id, orderin.book, orderin.amount, orderin.price, b.name, b.author, b.publisher, b.isbn,
	b.price, status, billref, bill.amount, bill.created FROM orderin
	LEFT OUTER JOIN book AS b ON book=b.id LEFT OUTER JOIN bill ON bill.id=billref`
	if book != nil {
		query += fmt.Sprintf(" WHERE book=%d", book.ID())
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Alert("Error occured while searching books: %s", err)
		return nil, err
	}
	defer rows.Close()
	ret = make([]OrderIn, 0, 30)
	idx := 0
	for ; rows.Next(); idx++ {
		ret = append(ret, OrderIn{})
		cur := &ret[idx]
		cur.Book = new(Book)
		bill := sql.NullInt64{}
		bamount := sql.NullFloat64{}
		bcreated := new(time.Time)
		var err = rows.Scan(&cur.id, &cur.Book.id, &cur.Amount, &cur.Price, &cur.Book.Name, &cur.Book.Author,
			&cur.Book.Publisher, &cur.Book.ISBN, &cur.Book.Price, &cur.Status, &bill, &bamount, &bcreated)
		if err != nil {
			log.Alert("Error when fetching row %d: %s", idx, err)
			return nil, err
		}
		if bill.Valid {
			cur.BillRef = new(Bill)
			id, err := bill.Value()
			if err != nil {
				log.Alert("Error when fetching row %d: %s", idx, err)
				return nil, err
			}
			cur.BillRef.id = id.(int32)
			amount, err := bamount.Value()
			if err != nil {
				log.Alert("Error when fetching row %d: %s", idx, err)
				return nil, err
			}
			cur.BillRef.Amount = amount.(float64)
			cur.BillRef.Created = *bcreated
		} else {
			cur.BillRef = nil
		}
	}
	return ret[:idx], nil
}

func GetOrderIn(id int32) (ret *OrderIn, err error) {
	res := db.QueryRow(`SELECT orderin.id, orderin.book, orderin.amount, orderin.price, b.name, b.author, b.publisher, b.isbn,
	b.price, status, billref, bill.amount, bill.created FROM orderin
	LEFT OUTER JOIN book AS b ON book=b.id LEFT OUTER JOIN bill ON bill.id=billref WHERE orderin.id=$1`, id)
	ret = new(OrderIn)
	ret.Book = new(Book)
	bill := sql.NullInt64{}
	bamount := sql.NullFloat64{}
	bcreated := new(time.Time)
	err = res.Scan(&ret.id, &ret.Book.id, &ret.Amount, &ret.Price, &ret.Book.Name, &ret.Book.Author,
		&ret.Book.Publisher, &ret.Book.ISBN, &ret.Book.Price, &ret.Status, &bill, &bamount, &bcreated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn("No such order: %d", id)
			return nil, nil
		}
		log.Alert("Error when fetching row: %s", err)
		return nil, err
	}
	if bill.Valid {
		ret.BillRef = new(Bill)
		id, err := bill.Value()
		if err != nil {
			log.Alert("Error when fetching row: %s", err)
			return nil, err
		}
		ret.BillRef.id = id.(int32)
		amount, err := bamount.Value()
		if err != nil {
			log.Alert("Error when fetching row: %s", err)
			return nil, err
		}
		ret.BillRef.Amount = amount.(float64)
		ret.BillRef.Created = *bcreated
	} else {
		ret.BillRef = nil
	}
	return
}

func GetOrderOut(book *Book) (ret []OrderIn, err error) {
	query := `SELECT orderout.id, orderout.book, amount, price, b.name, b.author, b.publisher, b.isbn,
	b.price, billref, bill.amount, bill.created FROM orderout
	LEFT OUTER JOIN book AS b ON book=b.id LEFT OUTER JOIN bill ON bill.id=billref`
	if book != nil {
		query += fmt.Sprintf(" WHERE book=%d", book.ID())
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Alert("Error occured while searching books: %s", err)
		return nil, err
	}
	defer rows.Close()
	ret = make([]OrderIn, 0, 30)
	idx := 0
	for ; rows.Next(); idx++ {
		ret = append(ret, OrderIn{})
		cur := &ret[idx]
		cur.Book = new(Book)
		cur.BillRef = new(Bill)
		var err = rows.Scan(&cur.id, &cur.Book.id, &cur.Amount, &cur.Price, &cur.Book.Name, &cur.Book.Author,
			&cur.Book.Publisher, &cur.Book.ISBN, &cur.Book.Price, &cur.BillRef.id, &cur.BillRef.Amount, &cur.BillRef.Created)
		if err != nil {
			log.Alert("Error when fetching row %d: %s", idx, err)
			return nil, err
		}
	}
	return ret[:idx], nil
}

/*

DB Creation SQL:

DROP TABLE IF EXISTS "orderin";
DROP TABLE IF EXISTS "orderout";

CREATE TABLE "orderin" (
	id serial PRIMARY KEY,
	book integer NOT NULL UNIQUE REFERENCES book,
	amount integer NOT NULL,
	price decimal(9,2) NOT NULL,
	status integer NOT NULL DEFAULT 0,
	billref integer NULL UNIQUE REFERENCES bill
);

CREATE TABLE "orderout" (
	id serial PRIMARY KEY,
	book integer NOT NULL UNIQUE REFERENCES book,
	amount integer NOT NULL,
	price decimal(9,2) NOT NULL,
	billref integer NOT NULL UNIQUE REFERENCES bill
);

*/
