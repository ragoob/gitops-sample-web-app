package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

}

func TestSumHandler(t *testing.T) {
	x := 5
	y := 6
	sum := x + y
	path := fmt.Sprintf("/calculate?x=%s&y=%s", strconv.Itoa(x), strconv.Itoa(y))
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SumHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	res, err := strconv.Atoi(strings.TrimSpace(rr.Body.String()))
	if err != nil {
		t.Error(err.Error())
	}
	if res != sum {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), sum)
	}

}
