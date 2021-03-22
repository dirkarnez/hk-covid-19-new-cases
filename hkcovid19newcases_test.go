package hkcovid19newcases

import (
	"testing"
	"fmt"
)

func TestShouldGet(t *testing.T) {
	fmt.Println("TestShouldGet")

	count, error := Get()

	fmt.Printf("count: %d\n", count)

	if error != nil {
		t.Errorf("Should not return error")
	}
}