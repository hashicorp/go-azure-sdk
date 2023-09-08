package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadIndustry string

const (
	PayloadIndustryIT                PayloadIndustry = "IT"
	PayloadIndustrybanking           PayloadIndustry = "Banking"
	PayloadIndustrybusinessServices  PayloadIndustry = "BusinessServices"
	PayloadIndustryconstruction      PayloadIndustry = "Construction"
	PayloadIndustryconsulting        PayloadIndustry = "Consulting"
	PayloadIndustryconsumerServices  PayloadIndustry = "ConsumerServices"
	PayloadIndustrycourierServices   PayloadIndustry = "CourierServices"
	PayloadIndustryeducation         PayloadIndustry = "Education"
	PayloadIndustryenergy            PayloadIndustry = "Energy"
	PayloadIndustryfinancialServices PayloadIndustry = "FinancialServices"
	PayloadIndustrygovernment        PayloadIndustry = "Government"
	PayloadIndustryhealthcare        PayloadIndustry = "Healthcare"
	PayloadIndustryhospitality       PayloadIndustry = "Hospitality"
	PayloadIndustryinsurance         PayloadIndustry = "Insurance"
	PayloadIndustrylegal             PayloadIndustry = "Legal"
	PayloadIndustrymanufacturing     PayloadIndustry = "Manufacturing"
	PayloadIndustryother             PayloadIndustry = "Other"
	PayloadIndustryrealEstate        PayloadIndustry = "RealEstate"
	PayloadIndustryretail            PayloadIndustry = "Retail"
	PayloadIndustrytelecom           PayloadIndustry = "Telecom"
	PayloadIndustryunknown           PayloadIndustry = "Unknown"
)

func PossibleValuesForPayloadIndustry() []string {
	return []string{
		string(PayloadIndustryIT),
		string(PayloadIndustrybanking),
		string(PayloadIndustrybusinessServices),
		string(PayloadIndustryconstruction),
		string(PayloadIndustryconsulting),
		string(PayloadIndustryconsumerServices),
		string(PayloadIndustrycourierServices),
		string(PayloadIndustryeducation),
		string(PayloadIndustryenergy),
		string(PayloadIndustryfinancialServices),
		string(PayloadIndustrygovernment),
		string(PayloadIndustryhealthcare),
		string(PayloadIndustryhospitality),
		string(PayloadIndustryinsurance),
		string(PayloadIndustrylegal),
		string(PayloadIndustrymanufacturing),
		string(PayloadIndustryother),
		string(PayloadIndustryrealEstate),
		string(PayloadIndustryretail),
		string(PayloadIndustrytelecom),
		string(PayloadIndustryunknown),
	}
}

func parsePayloadIndustry(input string) (*PayloadIndustry, error) {
	vals := map[string]PayloadIndustry{
		"it":                PayloadIndustryIT,
		"banking":           PayloadIndustrybanking,
		"businessservices":  PayloadIndustrybusinessServices,
		"construction":      PayloadIndustryconstruction,
		"consulting":        PayloadIndustryconsulting,
		"consumerservices":  PayloadIndustryconsumerServices,
		"courierservices":   PayloadIndustrycourierServices,
		"education":         PayloadIndustryeducation,
		"energy":            PayloadIndustryenergy,
		"financialservices": PayloadIndustryfinancialServices,
		"government":        PayloadIndustrygovernment,
		"healthcare":        PayloadIndustryhealthcare,
		"hospitality":       PayloadIndustryhospitality,
		"insurance":         PayloadIndustryinsurance,
		"legal":             PayloadIndustrylegal,
		"manufacturing":     PayloadIndustrymanufacturing,
		"other":             PayloadIndustryother,
		"realestate":        PayloadIndustryrealEstate,
		"retail":            PayloadIndustryretail,
		"telecom":           PayloadIndustrytelecom,
		"unknown":           PayloadIndustryunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadIndustry(input)
	return &out, nil
}
