package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/snowflakedb/gosnowflake"
)

func TestSnowConn() {
	account := ""
	user := ""
	pwd := ""
	database := ""
	warehouse := ""
	db, err := sql.Open("snowflake", user+":"+pwd+"@"+account+"/"+database+"?warehouse="+warehouse)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)

	for i := 0; i < 100; i++ {
		go func(j int) {
			tn := fmt.Sprintf("thread-%d", j)
			fmt.Println(tn)
			start := time.Now()
			read(db)
			fmt.Println(tn, ":", time.Since(start))
			done <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-done
	}

	defer db.Close()
}

func read(db *sql.DB) {
	rows, err := db.Query("SELECT system$wait(10);")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	defer rows.Close()
}
