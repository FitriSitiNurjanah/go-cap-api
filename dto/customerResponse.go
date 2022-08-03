package dto

type CustomerResponse struct {
<<<<<<< HEAD
	//`` berfungsi untuk penamaan dalam json di postman
	ID          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zip_code"`
=======
	ID          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zipcode"`
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
	DateOfBirth string `json:"date_of_birth" xml:"dateofbirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}
