package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type CustomerService interface {
<<<<<<< HEAD
	GetAllCustomer(string) (*[]dto.CustomerResponse, *errs.AppErr)
=======
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppErr)
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
<<<<<<< HEAD
	repository domain.CustomerRepository
}


func (s DefaultCustomerService) GetAllCustomer(status string) (*[]dto.CustomerResponse, *errs.AppErr){
	// // add proses here 
	cust, err := s.repository.FindAll(status)
	if err != nil{
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}

	var customerResponse []dto.CustomerResponse
	for _, data := range cust{
		customerResponse = append(customerResponse, data.ToDTO())
	}
	return &customerResponse, nil
}



func (s DefaultCustomerService) GetCustomerByID(customerID string)(*dto.CustomerResponse, *errs.AppErr){
	cust, err := s.repository.FindByID(customerID)
	if err != nil{
		return nil, err
	}
	response := cust.ToDTO()
	
	return &response, nil
}


func NewCustomerService(repository domain.CustomerRepository)DefaultCustomerService{
	return DefaultCustomerService{repository: repository}

}
=======
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppErr) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else if status == "" {
		status = ""
	} else {
		return nil, errs.NewBadRequestError("invalid request")
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var response []dto.CustomerResponse
	for _, customer := range customers {
		response = append(response, customer.ToDTO())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*dto.CustomerResponse, *errs.AppErr) {
	c, err := s.repo.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	response := c.ToDTO()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
