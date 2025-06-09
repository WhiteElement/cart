package db

import (
	"cartv2/cart/internal/item/item"
	"cartv2/cart/internal/shoppinglist/shoppinglist"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

func (db DB) QueryLastWrite() time.Time {
	rows, err := db.Conn.Query("SELECT * FROM public.\"Sync\"")
	if err != nil {
		return time.Time{}
	}

	var lastWrite time.Time
	if rows == nil {
		lastWrite = time.Now()
		db.Conn.Exec("INSERT INTO public.\"Sync\" (\"LastWrite\") VALUES ($1)", lastWrite)
	}

	rowsFound := 0
	for rows.Next() {
		rowsFound++
		if rowsFound > 1 {
			log.Printf("WARNING: More than one lastWrite Server Timestamp in Database found")
		}

		rows.Scan(&lastWrite)
	}

	return lastWrite
}

func (db DB) UpdateList(listId int, ts time.Time) error {
	_, err := db.Conn.Exec("UPDATE public.\"Lists\" SET \"Updated\" = $1 WHERE \"Id\" = $2", ts, listId)
	return err
}

func (db DB) QueryItem(id int) (item.Item, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE \"Items\".\"Id\" = '%d'", Items, id)
	row := db.Conn.QueryRow(sql)

	var item item.Item
	err := row.Scan(&item.Id, &item.Name, &item.Checked, &item.ListId, &item.Updated)

	return item, err
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
