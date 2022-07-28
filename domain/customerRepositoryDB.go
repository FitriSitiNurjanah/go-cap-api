package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //go get nama import ketika nama import error
)

type CustomerRepositoryDB struct{
	client *sqlx.DB
}

func NewCustomerRepositoryDB()CustomerRepositoryDB{
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	connStr := "postgres://postgres:fitri123@localhost/banking?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string)(*Customer, *errs.AppErr){
	query := "select * from customers where customer_id = $1"

	// row := d.client.QueryRow(query, customerID)
	// err := row.Scan(&c.ID, &c.Name,&c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)

	var c Customer

	err := d.client.Get(&c, query, customerID)
	if err != nil {	
		if err == sql.ErrNoRows{
			logger.Error("error customer data not found"+ err.Error())
			return nil, errs.NewNotFoundError("customer data not found")
		}else{
			logger.Error("error scanning customer data"+ err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		}
		return &c, nil
	}
		



func (d CustomerRepositoryDB) FindAll()([]Customer, error){

	query := "select * from customers"

	rows, err := d.client.Query(query)
	if err != nil {
		log.Println("error query data to customer table", err.Error())
		return nil, err
	} 

	var customers [] Customer
	for rows.Next(){

		var c Customer
		err := rows.Scan(&c.ID, &c.Name,&c.DateOfBirth, &c.City, &c.ZipCode, &c.ZipCode)
		if err!= nil{

			if err == sql.ErrNoRows{
				log.Println("error customer data not found", err.Error())
				return nil, errors.New("customer data not found")
			}else{
				log.Println("error scanning customer data", err.Error())
				return nil, err
			}
				
		}
		customers = append(customers, c)
	}
	return customers, nil
}