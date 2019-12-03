package main

import (
	"database/sql"
	"fmt"
	"log"

	clickhouse "github.com/ClickHouse/clickhouse-go"
)

func initClickhouse() {
	connect, err := sql.Open("clickhouse", "http://localhost:8123")
	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	// create initial tables
	_, err = connect.Exec(`
        CREATE TABLE IF NOT EXISTS file_metadata (
            owner character String,
            file_rows UInt8,
            file_size String,
            file_name String,
            file_uuid String,
            last_updated Date
        ) engine=Memory
    `)

	if err != nil {
		fmt.Printf("Error creating table 'file_metadata': [%s]", err)
	} else {
		fmt.Printf("Table 'file_metadata' created successfully")
	}

	_, err = connect.Exec(`
        CREATE TABLE IF NOT EXISTS user_actions (
            id Uint8,
            owner String,
            file_uuid String,
            action_taken Nested(
                action String,
                uuid String,
                historyState UInt8,
                column String,
                order String,
                value String,
                condition String,
                not UInt8
            ),
            action_timestamp timestamp without time zone
        ) engine=Memory
    `)

	if err != nil {
		fmt.Printf("Error creating table 'user_actions': [%s]", err)
	} else {
		fmt.Printf("Table 'user_actions' created successfully")
	}
}

func main() {
	initClickhouse()
}
