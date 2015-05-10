package main

import (
	. "github.com/jackyyf/gook/db"
	. "github.com/jackyyf/gook/models"
	"time"
)

func main() {
	Exec(`DROP TABLE IF EXISTS "orderin"`)
	Exec(`DROP TABLE IF EXISTS "orderout"`)
	Exec(`DROP TABLE IF EXISTS "bill"`)
	Exec(`DROP TABLE IF EXISTS "book"`)
	Exec(`DROP TABLE IF EXISTS "user"`)
	Exec(`CREATE TABLE "bill" (
	id serial PRIMARY KEY,
	amount decimal(9,2) NOT NULL,
	created timestamptz NOT NULL)`)
	Exec(`CREATE TABLE "user" (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL UNIQUE,
	pwd varchar(128),
	realname varchar(255) NOT NULL,
	gender integer,
	born timestamptz,
	admin boolean)`)
	Exec(`CREATE INDEX rname_idx ON "user"(realname)`)
	Exec(`CREATE TABLE "book" (
	id serial PRIMARY KEY,
	isbn char(13) NOT NULL UNIQUE,
	name varchar(255) NOT NULL,
	publisher varchar(255) NOT NULL,
	author varchar(255) NOT NULL,
	price decimal(9,2) NOT NULL,
	amount integer NOT NULL DEFAULT 0 CHECK(amount>=0))`)
	Exec(`CREATE INDEX bname_idx ON "book"(name)`)
	Exec(`CREATE INDEX bpublisher_idx ON "book"(publisher)`)
	Exec(`CREATE INDEX bauthor ON "book"(author)`)
	Exec(`CREATE TABLE "orderin" (
	id serial PRIMARY KEY,
	book integer NOT NULL REFERENCES book,
	amount integer NOT NULL,
	price decimal(9,2) NOT NULL,
	status integer NOT NULL DEFAULT 0,
	billref integer NULL UNIQUE REFERENCES bill)`)
	Exec(`CREATE TABLE "orderout" (
	id serial PRIMARY KEY,
	book integer NOT NULL REFERENCES book,
	amount integer NOT NULL,
	price decimal(9,2) NOT NULL,
	billref integer NOT NULL UNIQUE REFERENCES bill)`)
	user := new(User)
	user.SetPassword("admin")
	user.SetAdmin()
	user.Name = "admin"
	user.RealName = "超级管理员"
	user.Gender = 1
	user.Born = time.Now()
	user.Create()
}
