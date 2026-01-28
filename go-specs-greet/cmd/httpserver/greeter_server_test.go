package main_test

import (
	"testing"

	"github.com/quii/go-specs-greet/specifications"
	go_specs_greet "github.com/tomhockett/learn-go-with-tests/go-specs-greet"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
