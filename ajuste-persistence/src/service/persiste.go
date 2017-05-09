package service

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	log.Print("Estabelecendo conexao")

	ldb, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/teste")
	if err != nil {
		log.Panic("Falha ao estabelecer conexao", err)
	}
	ldb.SetMaxOpenConns(20)
	ldb.SetMaxIdleConns(10)
	ldb.SetConnMaxLifetime(1 * time.Minute)

	if err != nil {
		log.Panicf("conexao falhou: %v", err)
	}

	log.Printf("Conexao Estabelecida!")
	db = ldb
}

func Persist(requests ...Request) {

	l := time.Now()
	query := "INSERT INTO TBSETR_FINANCE_ADJUSTMENT_REQ VALUES"

	vet := []interface{}{}
	for _, entry := range requests {

		query += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"

		vet = append(vet, entry.Id, entry.Type, entry.Date, 1,
			entry.SettlementAdjustment, entry.GrossValue, entry.UserId, entry.TechnologyType, 1,
			entry.Customer, entry.ModCustomer, entry.BatchId, entry.MovementType, entry.RefundStatus, entry.AdjustmentReason,
			entry.AdjustmentComments, 1, entry.ProductCode, 1,
		)
	}

	log.Printf("Persistindo")

	query = query[:len(query)-1]

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	result, err := stmt.Exec(vet...)
	log.Printf("Persistido em  %v seg", time.Since(l).Seconds())
	log.Printf("Result %v %v", result, err)
}
