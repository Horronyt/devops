package main

	import (
		"database/sql"
		"fmt"
		"log"
		"net/http"

		_ "github.com/lib/pq"
	)

	const (
		dbUser = "user"
		dbPassword = "password"
		dbName = "mydb"
	)

	func main() {
		connStr := fmt.Sprintf("host=db user =%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			rows, err := db.Query("SELECT id, name FROM mytable")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var id int
				var name string
				if err := rows.Scan(&id, &name); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "%d: %s\n", id, name)
			}
		})

		log.Println("Loading server on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
