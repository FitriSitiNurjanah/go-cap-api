package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
<<<<<<< HEAD
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
		

	
func (d CustomerRepositoryDB) FindAll(status string)([]Customer, *errs.AppErr){
	query := "select * from customers"
=======
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{client}
}

func (s CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppErr) {

	var query string
	var err error

>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
	var customers []Customer

	if status == "" {
		query = "select * from customers"
<<<<<<< HEAD
		err := d.client.Select(&customers, query)
		if err!= nil{
				log.Println("error scanning customer data"+ err.Error())
				return nil, errs.NewUnexpectedError("unexpected database error")
			}	
	}else if status == "active" || status == "inactive"{
		if status == "active" {
			status = "1"
			}else {
			status = "0"
			}

		query = "select * from customers where status = $1"
		err := d.client.Select(&customers, query, status)
		if err!= nil{
				log.Println("error scanning customer data"+ err.Error())
				return nil, errs.NewUnexpectedError("unexpected database error")
		}	
	}
	return customers, nil
}

=======
		err = s.db.Select(&customers, query)
	} else {
		query = "select * from customers where status = $1"
		err = s.db.Select(&customers, query, status)
	}

	if err != nil {
		logger.Error("error fetch data to customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil

}

func (s CustomerRepositoryDB) FindByID(id string) (*Customer, *errs.AppErr) {

	query := "select * from customers where customer_id = $1"

	var c Customer

	err := s.db.Get(&c, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("error scanning customer data ", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

	}

	return &c, nil

}
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
