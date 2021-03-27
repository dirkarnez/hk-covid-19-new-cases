package hkcovid19

import (
	"strings"
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

func TestShouldGetCommunityTestingCentres(t *testing.T) {
	fmt.Println("GetCommunityTestingCentres")

	centres, err := GetCommunityTestingCentres()
	if err != nil {
		t.Errorf("Should not return error")
	}

	fmt.Printf("count of centres: %d\n", len(centres))
	fmt.Printf("%s\n", strings.Join(centres, ",\n"))
}
