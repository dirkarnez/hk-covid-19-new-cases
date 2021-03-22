package hkcovid19newcases

import (
	"testing"
	"fmt"
)

func TestShouldGet(t *testing.T) {
	fmt.Println("TestShouldGet")

	count, err := Get()
	if err != nil {
		t.Errorf("Should not return error")
	}

	fmt.Printf("count: %d\n", count)
}