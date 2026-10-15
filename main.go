package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// _ "modernc.org/sqlite"
)

type Phone struct {
	id     int
	vendor string
	model  string
	price  sql.NullInt64
}

func main() {
	const pgConfig = "user=postgres password=postgres dbname=mydb sslmode=disable port=5432"
	db, err := sql.Open("postgres", pgConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	query := fmt.Sprintf(`
		INSERT INTO phones (vendor, model, price)
		VALUES ('%s', '%s', %d);`, "Huawey", "Simple 1", 1000)

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query(`SELECT * FROM phones;`)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	phones := []Phone{}
	for rows.Next() {
		p := Phone{}
		err := rows.Scan(&p.id, &p.vendor, &p.model, &p.price)
		if err != nil {
			log.Println(err)
			continue
		}
		if !p.price.Valid {
			p.price.Int64 = 0
		}
		phones = append(phones, p)
	}
	for _, v := range phones {
		fmt.Println(v.vendor, v.model, v.price.Int64)
	}

}
