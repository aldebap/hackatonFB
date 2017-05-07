package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//?multiStatements=true
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/teste")
	if err != nil {
		log.Panicf("conexao falhou: %v", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1 * time.Minute)

	query := "INSERT INTO TBSETR_FINANCE_ADJUSTMENT_REQ(NU_REQUEST, CD_REQUEST_TYPE, NU_CUSTOMER, NU_MOD_CUSTOMER, NU_BATCH, DH_REQUEST) VALUES"

	if err != nil {
		log.Panicf("Falhou statement: %v", err)
	}

	vet := []interface{}{}
	l := time.Now()
	for i := 0; i < 10000; i++ {
		query += "(?, ?, ?, ?, ?, ?),"

		vet = append(vet, i, 1, i, i, 1, time.Now())
	}

	log.Printf("Persistindo")

	query = query[:len(query)-1]

	stmt, err := db.Prepare(query)

	result, err := stmt.Exec(vet...)
	log.Printf(">>> %v", time.Since(l).Seconds())

	log.Printf("Result %v %v", result, err)
}
