package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/adityapwr/banking/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
		rows, err = d.client.Query(findAllSql)

	} else {
		findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers where status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		log.Println("Error while fetching rows" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.Pincode, &c.Status)
		if err != nil {
			log.Println("Error while scanning rows" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected scan error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.Pincode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning rows" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
