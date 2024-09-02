package main_test

import (
	"testing"

	"github.com/ArsaGit/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	specifications.GreetSpecification(t, nil)
}
