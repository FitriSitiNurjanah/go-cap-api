package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "User1", "Jakarta", "12345", "2022-01-01", "1"},
<<<<<<< HEAD
		{"2", "User2", "Surabaya", "67890", "2022-02-02", "1"},
		{"3", "User3", "Bandung", "4556", "2022-03-03", "1"},
		{"4", "User4", "Bogor", "7894", "2022-04-04", "1"},
	}
	return CustomerRepositoryStub{Customer: customers}
}
=======
		{"2", "User2", "Surabaya", "67890", "2022-01-01", "1"},
		{"3", "User3", "Jakarta", "12345", "2022-01-01", "1"},
		{"4", "User4", "Surabaya", "67890", "2022-01-01", "1"},
	}

	return CustomerRepositoryStub{Customer: customers}
}
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
