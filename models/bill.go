package models

import (
	"database/sql"
	"errors"
	"github.com/jackyyf/gook/db"
	"github.com/jackyyf/gook/utils/log"
	"time"
)

type Bill struct {
	id      int32
	Amount  float64
	Created time.Time
}

func (b *Bill) ID() int32 {
	return b.id
}

func (b *Bill) Create(tx *sql.Tx) (err error) {
	if b == nil {
		err = errors.New("Error: unable to create nil bill.")
		log.Alert(err.Error())
		return
	}
	b.Created = time.Now()
	err = tx.QueryRow(
		`INSERT INTO "bill" (amount, created) VALUES ($1, $2) RETURNING id`,
		b.Amount, b.Created).Scan(&b.id)
	if err != nil {
		log.Error("Create bill failed: %s", err)
		return
	}
	log.Info("New bill(ID=%d) created successfully.", b.id)
	return
}

func GetBillsAfter(t time.Time) (bills []Bill, err error) {
	rows, err := db.Query(
		`SELECT id, amount, created, "in". FROM "bill"	WHERE created >= $1`, t)
	if err != nil {
		log.Alert("Error occured when searching bills: %s", err)
		return
	}
	defer rows.Close()
	bills = make([]Bill, 0, 30)
	idx := 0
	for ; rows.Next(); idx++ {
		bills = append(bills, Bill{})
		cur := &bills[idx]
		err = rows.Scan(&cur.id, &cur.Amount, &cur.Created)
		if err != nil {
			log.Alert("Error when fetching row %d: %s", idx, err)
			return nil, err
		}
	}
	return bills[:idx], nil
}

func GetSummaryAfter(t time.Time) (in, out float64, err error) {
	rows, err := db.Query(`SELECT sum(amount) AS total FROM "bill"
		WHERE created >= $1 GROUP BY amount > 0`, t)
	if err != nil {
		log.Alert("Error occured when calculating bill summary: %s", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		sum := 0
		err = rows.Scan(&sum)
		if err != nil {
			log.Alert("Error when fetching row: %s", err)
			return 0., 0., err
		}
		if sum < 0 {
			out = float64(-sum)
		} else if sum > 0 {
			in = float64(sum)
		}
	}
	return
}

/*

DB Creation SQL:

DROP TABLE IF EXISTS "bill";

CREATE TABLE "bill" (
	id serial PRIMARY KEY,
	amount decimal(9,2) NOT NULL,
	created timestamptz NOT NULL
);

*/
