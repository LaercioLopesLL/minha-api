package main

import (
	"net/http"
	"testing"
)

func TestHoraAtual(t *testing.T) {
	res, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code errado: %d", res.StatusCode)
	}
}