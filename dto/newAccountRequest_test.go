package dto

import (
	"net/http"
	"testing"
)

func Test_validate_min_account_balance_new_account(t *testing.T) {
	//AAA
	//Arrange
	newAccountRequest := NewAccountRequest{
		CustomerId:  "123",
		AccountType: "saving",
		Amount:      3000,
	}

	//Act
	err := newAccountRequest.Validate()

	//Assert
	if err == nil {
		t.Error("Expected error but got nil")
	}
	if err.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500 but got", err.Code)
	}
}
