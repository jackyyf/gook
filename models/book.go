package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackyyf/gook/db"
	"github.com/jackyyf/gook/utils/isbn"
	"github.com/jackyyf/gook/utils/log"
	"strings"
	"unicode/utf8"
)

type Book struct {
	id        int32
	ISBN      string
	Name      string
	Publisher string
	Author    string
	Price     float64
	Amount    int32
}

func (me *Book) ID() int32 {
	return me.id
}

func (me *Book) Create() (err error) {
	if me == nil {
		err = errors.New("Unable to create nil book.")
		log.Alert(err.Error())
		return
	}
	err = db.QueryRow(
		`INSERT INTO "book" (isbn, name, publisher, author, price, amount)
		VALUES ($1, $2, $3, $4, $5, 0) RETURNING id`,
		me.ISBN, me.Name, me.Publisher, me.Author, me.Price).Scan(&(me.id))
	if err != nil {
		log.Alert("Create new book failed: %s", err)
		return
	}

	log.Info("New book(ID=%d) created successfully.", me.id)
	return nil
}

func (me *Book) Save() (err error) {
	if me == nil {
		err = errors.New("Unable to create nil book.")
		log.Alert(err.Error())
		return
	}
	res, err := db.Exec(
		`UPDATE "book" SET isbn=$1, name=$2, publisher=$3, author=$4, price=$5, amount=$6 WHERE id=$7`,
		me.ISBN, me.Name, me.Publisher, me.Author, me.Price, me.Amount, me.id)
	if err != nil {
		log.Error("Commit book changes failed: %s", err)
		return err
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting chagnes to an unexist user!")
		log.Alert(err.Error())
		return err
	} else {
		log.Info("Commit changes for book(ID=%d) successfully.", me.id)
		return nil
	}
}

func (me *Book) Delete() (err error) {
	if me == nil {
		err = errors.New("Error: unable to delete nil book.")
		log.Alert("%s", err)
		return
	}
	res, err := db.Exec(`DELETE FROM "book" WHERE id=$1`, me.id)
	if err != nil {
		log.Error("Delete book(ID=%d) failed: %s", me.id, err)
		return err
	} else if t, _ := res.RowsAffected(); t == 0 {
		log.Warn("Trying to delete an unexist book(ID=%d)", me.id)
	} else {
		log.Info("Book(ID=%d) deleted successfully.", me.id)
	}
	return nil
}

func GetBook(id int32) (ret *Book, err error) {
	row := db.QueryRow(
		`SELECT id, isbn, name, publisher, author, price, amount FROM "book" WHERE id=$1 LIMIT 1`,
		id)
	ret = new(Book)
	if err := row.Scan(&ret.id, &ret.ISBN, &ret.Name, &ret.Publisher, &ret.Author,
		&ret.Price, &ret.Amount); err != nil {
		if err == sql.ErrNoRows {
			log.Warn("Book(ID=%d) not exists.", id)
			return nil, nil
		} else {
			log.Alert("Unknown error when fetching book(ID=%d): %s", id, err)
			return nil, err
		}
	}
	log.Info("Book(ID=%d) fetched successfully.", id)
	return
}

func SearchBooks(isbn_str string, names []string, publishers []string, authors []string, offset, length int) (ret []Book, err error) {
	query := `SELECT id, isbn, name, publisher, author, price, amount FROM "book"`
	isbn_str = isbn.Normalize(isbn_str)
	where_phase := ""
	arg_list := make([]interface{}, 0, len(names)+2)
	arg_num := 0
	if isbn_str != "" {
		where_phase = " WHERE isbn = $1"
		arg_list = append(arg_list, isbn_str)
		arg_num++
	} else {
		if names != nil {
			subphases := make([]string, 0, len(names))
			for _, keyword := range names {
				if utf8.RuneCountInString(keyword) <= 2 {
					continue
				}
				arg_num++
				subphase := fmt.Sprintf("name LIKE $%d", arg_num)
				subphases = append(subphases, subphase)
				arg_list = append(arg_list, "%"+keyword+"%")
			}
			if len(subphases) > 0 {
				where_phase = " WHERE (" + strings.Join(subphases, " OR ") + ")"
			}
		}
		if publishers != nil {
			subphases := make([]string, 0, len(publishers))
			for _, keyword := range publishers {
				if utf8.RuneCountInString(keyword) <= 2 {
					continue
				}
				arg_num++
				subphase := fmt.Sprintf("publisher LIKE $%d", arg_num)
				subphases = append(subphases, subphase)
				arg_list = append(arg_list, "%"+keyword+"%")
			}
			if len(subphases) > 0 {
				if where_phase != "" {
					where_phase = " WHERE ("
				} else {
					where_phase += " AND ("
				}
				where_phase += strings.Join(subphases, " OR ") + ")"
			}
		}
		if authors != nil {
			subphases := make([]string, 0, len(authors))
			for _, keyword := range authors {
				if utf8.RuneCountInString(keyword) <= 2 {
					continue
				}
				arg_num++
				subphase := fmt.Sprintf("author LIKE $%d", arg_num)
				subphases = append(subphases, subphase)
				arg_list = append(arg_list, "%"+keyword+"%")
			}
			if len(subphases) > 0 {
				if where_phase != "" {
					where_phase = " WHERE ("
				} else {
					where_phase += " AND ("
				}
				where_phase += strings.Join(subphases, " OR ") + ")"
			}
		}
	}
	limit_phase := " ORDER BY \"id\" DESC"
	if offset >= 0 {
		limit_phase += fmt.Sprintf(" OFFSET %d", offset)
	}
	if length > 0 {
		limit_phase += fmt.Sprintf(" LIMIT %d", length)
	}
	query += where_phase + limit_phase
	rows, err := db.Query(query, arg_list...)
	if err != nil {
		log.Alert("Error occured while searching books: %s", err)
		return nil, err
	}
	defer rows.Close()
	if length > 0 {
		ret = make([]Book, 0, length)
	} else {
		ret = make([]Book, 0, 30)
	}
	idx := 0
	for ; rows.Next(); idx++ {
		ret = append(ret, Book{})
		cur := &ret[idx]
		err = rows.Scan(&cur.id, &cur.ISBN, &cur.Name, &cur.Publisher,
			&cur.Author, &cur.Price, &cur.Amount)
		if err != nil {
			log.Alert("Error when fetching row %d: %s", idx, err)
			return nil, err
		}
	}
	return ret[:idx], nil
}

/*

DB Creation SQL:

DROP TABLE IF EXISTS "book";

CREATE TABLE "book" (
	id serial PRIMARY KEY,
	isbn char(13) NOT NULL UNIQUE,
	name varchar(255) NOT NULL,
	publisher varchar(255) NOT NULL,
	author varchar(255) NOT NULL,
	price decimal(9,2) NOT NULL,
	amount integer NOT NULL DEFAULT 0
);

CREATE INDEX bname_idx ON "book"(name);
CREATE INDEX bpublisher_idx ON "book"(publisher);
CREATE INDEX bauthor ON "book"(author);

*/
