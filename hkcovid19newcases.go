package hkcovid19

import (
	"github.com/antchfx/jsonquery"
	"strconv"
)

func GetNewCases() (uint64, error)  {
	doc, err := jsonquery.LoadURL("https://chp-dashboard.geodata.gov.hk/covid-19/data/keynum.json")
	if err != nil {
		return 0, err
	}

	confirmedNode := jsonquery.FindOne(doc, "Confirmed")
	pastConfirmedNode := jsonquery.FindOne(doc, "P_Confirmed")

	var confirmed uint64 = 0
	var pastConfirmed uint64 = 0

	if confirmedNode != nil {
		confirmed, err = strconv.ParseUint(confirmedNode.InnerText(), 10, 64)
		if err != nil {
			return 0, err
		}
	}

	if pastConfirmedNode != nil {
		pastConfirmed, err = strconv.ParseUint(pastConfirmedNode.InnerText(), 10, 64)
		if err != nil {
			return 0, err
		}
	}


	return confirmed - pastConfirmed, nil;
}
