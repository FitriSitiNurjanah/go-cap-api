package domain

import (
	"capi/dto"
	"capi/errs"
)

type Customer struct {
<<<<<<< HEAD
	//`` berfungsi untuk penamaan dalam json di postman
	ID          string `json:"id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status" `
}

//I untuk interface didepan nama
type CustomerRepository interface {
=======
	ID          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	//  status -> "1", "0", ""
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
	FindAll(string) ([]Customer, *errs.AppErr)
	FindByID(string) (*Customer, *errs.AppErr)
}

<<<<<<< HEAD
func (c Customer) convertStatusName()string{
	statusName := "active"
	if c.Status == "0"{
=======
func (c Customer) convertStatusName() string {
	statusName := "active"
	if c.Status == "0" {
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
		statusName = "inactive"
	}
	return statusName
}

<<<<<<< HEAD

func (c Customer) ToDTO() dto.CustomerResponse{
	
	return dto.CustomerResponse{
		ID: c.ID,
		Name: c.Name,
		DateOfBirth: c.DateOfBirth,
		City: c.City,
		ZipCode: c.ZipCode,
		Status: c.convertStatusName(),
=======
func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
		City:        c.City,
		ZipCode:     c.ZipCode,
		Status:      c.convertStatusName(),
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
	}
}
