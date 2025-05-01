package db

import (
	"cartv2/cart/internal/item/item"
	"cartv2/cart/internal/shoppinglist/shoppinglist"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

type DB struct {
	Conn *sql.DB
}

const (
	Items string = "public.\"Items\""
	Lists string = "public.\"Lists\""
)

func connString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("PGDBUSER"), os.Getenv("PGDBPASS"), os.Getenv("PGDBHOST"), os.Getenv("PGDBPORT"), os.Getenv("PGDB"))
}

func NewConn() DB {
	connString := connString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Could not connect to database:\nconnString: '%s'\nDatabase: '+%v'", connString, db)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return DB{Conn: db}
}

func (db DB) QueryAllItems() []item.Item {
	rows, err := db.Conn.Query(fmt.Sprintf("SELECT * FROM %s", Items))
	var items []item.Item
	if err != nil {
		return items
	}

	for rows.Next() {
		var item item.Item
		rows.Scan(&item.Id, &item.Name, &item.Checked, &item.ListId, &item.Updated)
		items = append(items, item)
	}

	return items
}

func (db DB) QueryItemsFromList(id int) []item.Item {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE \"Items\".\"ListId\" = '%d'", Items, id)
	rows, err := db.Conn.Query(sql)
	var items []item.Item
	if err != nil {
		log.Printf("Error querying items from list: id '%d', err '%s'", id, err.Error())
		return items
	}

	for rows.Next() {
		var item item.Item
		rows.Scan(&item.Id, &item.Name, &item.Checked, &item.ListId, &item.Updated)
		items = append(items, item)
	}

	return items
}

func (db DB) QueryAllLists() ([]shoppinglist.List, error) {
	rows, err := db.Conn.Query(fmt.Sprintf("SELECT * FROM %s", Lists))
	var lists []shoppinglist.List
	if err != nil {
		return lists, err
	}

	for rows.Next() {
		var list shoppinglist.List
		err = rows.Scan(&list.Id, &list.Name, &list.Archived, &list.Created, &list.Updated)
		lists = append(lists, list)
	}

	if err != nil {
		return lists, err
	}

	return lists, nil
}

func (db DB) QueryList(id int) shoppinglist.List {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE \"Lists\".\"Id\" = '%d'", Lists, id)
	row := db.Conn.QueryRow(sql)

	var list shoppinglist.List
	row.Scan(&list.Id, &list.Name, &list.Archived, &list.Created, &list.Updated)

	return list
}

func format(collection []string, surroundings string) string {
	var sb strings.Builder

	const SEP = ", "
	for _, c := range collection {
		_, _ = sb.WriteString(surroundings)
		_, _ = sb.WriteString(c)
		_, _ = sb.WriteString(surroundings)
		_, _ = sb.WriteString(SEP)
	}

	return sb.String()[:sb.Len()-len(SEP)]
}
