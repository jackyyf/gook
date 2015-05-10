package models

import (
	"database/sql"
	"errors"
	"github.com/jackyyf/gook/db"
	"github.com/jackyyf/gook/utils/log"
	"github.com/jackyyf/gook/utils/password"
	"time"
)

type User struct {
	// NEVER change id of an user object.
	id         int32
	hashed_pwd string
	Name       string
	RealName   string
	Sex        bool
	Born       time.Time
	admin      bool
}

func (me *User) ID() int32 {
	return me.id
}

func (me *User) SetPassword(newpwd string) {
	me.hashed_pwd = password.Generate(newpwd)
}

func (me *User) Login(pwd string) bool {
	if !password.Check(pwd, me.hashed_pwd) {
		return false
	}
	return true
}

func (me *User) IsAdmin() bool {
	return me.admin
}

func (me *User) SetAdmin() {
	me.admin = true
}

func (me *User) ClearAdmin() {
	me.admin = false
}

func (me *User) Create() (err error) {
	if me == nil {
		err = errors.New("Error: unable to create nil user.")
		log.Alert(err.Error())
		return
	}
	if me.hashed_pwd == "" {
		err = errors.New("No password for new user.")
		log.Alert(err.Error())
		return
	}
	err = db.QueryRow(
		`INSERT INTO "user" (name, pwd, realname, sex, born, admin) VALUES
		($1, $2, $3, $4, $5, FALSE) RETURNING id`,
		me.Name, me.hashed_pwd, me.RealName, me.Sex, me.Born).Scan(&(me.id))
	if err != nil {
		log.Error("Create user failed: %s", err)
		return
	}
	log.Info("New user(ID=%d) created successfully.", me.id)
	return
}

func (me *User) Save() (err error) {
	if me == nil {
		err = errors.New("Error: unable to save nil user.")
		log.Alert("%s", err)
		return
	}
	res, err := db.Exec(
		`UPDATE "user" SET name=$1, pwd=$2, realname=$3, sex=$4, born=$5, admin=$6 WHERE id=$7`,
		me.Name, me.hashed_pwd, me.RealName, me.Sex, me.Born, me.admin, me.id)
	if err != nil {
		log.Error("Commit user changes failed: %s", err)
		return err
	} else if t, _ := res.RowsAffected(); t == 0 {
		err = errors.New("Commiting chagnes to an unexist user!")
		log.Alert(err.Error())
		return
	} else {
		log.Info("Commit changes for user(ID=%d) successfully.", me.id)
		return nil
	}
}

func (me *User) Delete() (err error) {
	if me == nil {
		err = errors.New("Error: unable to delete nil user.")
		log.Alert("%s", err)
		return
	}
	res, err := db.Exec(`DELETE FROM "user" WHERE id=$1`, me.id)
	if err != nil {
		log.Error("Delete user(ID=%d) failed: %s", me.id, err)
		return err
	} else if t, _ := res.RowsAffected(); t == 0 {
		log.Warn("Trying to delete an unexist user(ID=%d)", me.id)
	} else {
		log.Info("User(ID=%d) deleted successfully.", me.id)
	}
	return nil
}

func GetUser(id int32) (ret *User, err error) {
	row := db.QueryRow(
		`SELECT id, name, pwd, realname, sex, born, admin FROM "user" WHERE id=$1 LIMIT 1`,
		id)
	ret = new(User)
	if err := row.Scan(&ret.id, &ret.Name, &ret.hashed_pwd, &ret.RealName, &ret.Sex,
		&ret.Born, &ret.admin); err != nil {
		if err == sql.ErrNoRows {
			log.Warn("User(ID=%d) not exists.", id)
			return nil, nil
		} else {
			log.Alert("Unknown error when fetching user(ID=%d): %s", id, err)
			return nil, err
		}
	}
	log.Info("User(ID=%d) fetched successfully.", id)
	return
}

func GetUserByName(name string) (ret *User, err error) {
	row := db.QueryRow(
		`SELECT id, name, pwd, realname, sex, born, admin FROM "user" WHERE name=$1 LIMIT 1`,
		name)
	ret = new(User)
	if err := row.Scan(&ret.id, &ret.Name, &ret.hashed_pwd, &ret.RealName, &ret.Sex,
		&ret.Born, &ret.admin); err != nil {
		if err == sql.ErrNoRows {
			log.Warn("User(name=%s) not exists.", name)
			return nil, nil
		} else {
			log.Alert("Unknown error when fetching user(name=%s): %s", name, err)
			return nil, err
		}
	}
	log.Info("User(name=%d) fetched successfully.", name)
	return
}

func DeleteUsers(IDs ...int32) (success_ids []int32) {
	// Do note: nil means database error, while empty slice means no successful delete.
	if len(IDs) == 0 {
		return IDs
	}
	tx, err := db.Transaction()
	if err != nil {
		log.Alert("Delete users failed while creating transcation: %s", err)
		return nil
	}
	stmt, err := tx.Prepare(`DELETE FROM "user" WHERE id=$1`)
	if err != nil {
		log.Alert("Delete users failed while creating prepared statement: %s", err)
		return nil
	}
	success_ids = make([]int32, 0, len(IDs))
	for _, id := range IDs {
		res, err := stmt.Exec(id)
		if err != nil {
			log.Alert("Delete user(ID=%d) failed: %s", id, err)
			log.Alert("Rollback.")
			stmt.Close()
			tx.Rollback()
			return nil
		}
		if t, _ := res.RowsAffected(); t == 0 {
			log.Warn("Deleting a not exist user(ID=%d), skipped.", id)
		} else {
			success_ids = append(success_ids, id)
		}
	}
	// Ignore close error
	stmt.Close()
	err = tx.Commit()
	if err != nil {
		log.Alert("Delete users failed while commiting transcation: %s", err)
		log.Alert("Rollback.")
		tx.Rollback()
		return nil
	}
	log.Info("OK, %d user(s) deleted", len(success_ids))
	return
}

/*

DB Creation SQL:

DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL UNIQUE,
	pwd varchar(128),
	realname varchar(255) NOT NULL,
	sex boolean,
	born timestamptz,
	admin boolean
);

CREATE INDEX rname_idx ON "user"(realname);

*/
