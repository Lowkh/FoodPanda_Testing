package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func TestHandlerAlive(t *testing.T) {
	req, err := http.NewRequest("GET","localhost:8080/feature2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feature1)

	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := "Feature1"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestServerRouting(t * testing.T){
	server := httptest.NewServer(handlers())
	defer server.Close()
	result, err := http.Get(fmt.Sprintf("%s/feature1", server.URL))

	if err != nil{
		t.Fatalf("Get Request error: %v", err)
	}
	defer result.Body.Close()
	if result.StatusCode != http.StatusOK{
		t.Errorf("Expected status OK but got %v", result.Status)
	}

}