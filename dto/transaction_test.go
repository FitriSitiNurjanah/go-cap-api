package dto

import (
	"capi/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTranaction_valid_transaction(t *testing.T) {
	tx := TransactionRequest{
		TransactionType: WITHDRAWAL,
		Amount: 1,
	}

	res := tx.Validate()

	assert.Nil(t, res)
}

func TestValidateTransaction_Invalid_transaction_type(t *testing.T){
	tx := TransactionRequest{
		Amount: 1,
		TransactionType: "Invalid Type",
	}

	res := tx.Validate()

	expect := errs.NewValidationError("")

	assert.NotNil(t, res)
	assert.Equal(t, expect.Code, res.Code)
}

func TestValidateTransaction_invalid_transaction_amount(t *testing.T){
	tx := TransactionRequest{
		TransactionType: WITHDRAWAL,
		Amount: -1,
	}
	res := tx.Validate()

	expect := errs.NewValidationError("")

	assert.NotNil(t, res)
	assert.Equal(t, expect.Code, res.Code)
}