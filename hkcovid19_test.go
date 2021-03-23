package hkcovid19

import (
	"testing"
	"fmt"
)

func TestShouldGetNewCases(t *testing.T) {
	fmt.Println("TestShouldGetNewCases")

	count, err := GetNewCases()
	if err != nil {
		t.Errorf("Should not return error")
	}

	fmt.Printf("new cases: %d\n", count)
}
