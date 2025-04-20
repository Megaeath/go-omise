package chargeclient

import (
	"go-cli/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendCharge(t *testing.T) {
	// Mock server to simulate the API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if r.URL.Path != "/api/charge" {
			t.Errorf("Expected URL path '/api/charge', got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	// Test data
	row := model.DonationRow{
		Name:           "John Doe",
		AmountSubunits: 100,
	}
	host := mockServer.URL

	// Call the function
	success, err := SendCharge(row, host)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !success {
		t.Fatalf("Expected success to be true, got false")
	}
}

func TestSendCharge_Failure(t *testing.T) {
	// Mock server to simulate a failure response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	// Test data
	row := model.DonationRow{
		Name:           "Jane Doe",
		AmountSubunits: 200,
	}
	host := mockServer.URL

	// Call the function
	success, err := SendCharge(row, host)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	if success {
		t.Fatalf("Expected success to be false, got true")
	}
}
