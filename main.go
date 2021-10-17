package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	posts, err := getPosts(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("ID: %v, name: %s\n%s\n%s\n\n", post.id, post.name, post.content, post.createdAt.String())
	}
}

type post struct {
	id        int
	name      string
	content   string
	createdAt time.Time
}

func getPosts(db *sql.DB) ([]post, error) {
	query := "SELECT id, name, content, created_at from posts order by id"

	posts := make([]post, 0)

	rows, err := db.Query(query)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var p post

		err := rows.Scan(&p.id, &p.name, &p.content, &p.createdAt)
		if err != nil {
			return posts, err
		}

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}
