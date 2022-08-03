package app

import (
	"capi/dto"
	"capi/mocks/service"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var mockService *service.MockCustomerService
var ch CustomerHandlers
var router *mux.Router

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	return func() {
		router = nil
		defer ctrl.Finish()
	}

}


func Test_Should_return_customer_with_status_code_200(t *testing.T){
	//arange 
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)

	dummyCustomer := []dto.CustomerResponse{
		{"1", "User1", "Jakarta", "12345", "2022-01-01", "1"},
		{"2", "User2", "Surabaya", "67890", "2022-01-01", "1"},
	}

	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomer, nil)
	ch := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//assert
	fmt.Println(recorder.Body.String())
	if recorder.Code != http.StatusOK{
		t.Error("failed testing status code")
	}
}
