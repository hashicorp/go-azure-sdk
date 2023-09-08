package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadBrand string

const (
	PayloadBrandadobe           PayloadBrand = "Adobe"
	PayloadBrandamericanExpress PayloadBrand = "AmericanExpress"
	PayloadBrandcapitalOne      PayloadBrand = "CapitalOne"
	PayloadBranddhl             PayloadBrand = "Dhl"
	PayloadBranddocuSign        PayloadBrand = "DocuSign"
	PayloadBranddropbox         PayloadBrand = "Dropbox"
	PayloadBrandfacebook        PayloadBrand = "Facebook"
	PayloadBrandfirstAmerican   PayloadBrand = "FirstAmerican"
	PayloadBrandmicrosoft       PayloadBrand = "Microsoft"
	PayloadBrandnetflix         PayloadBrand = "Netflix"
	PayloadBrandother           PayloadBrand = "Other"
	PayloadBrandscotiabank      PayloadBrand = "Scotiabank"
	PayloadBrandsendGrid        PayloadBrand = "SendGrid"
	PayloadBrandstewartTitle    PayloadBrand = "StewartTitle"
	PayloadBrandsyrinxCloud     PayloadBrand = "SyrinxCloud"
	PayloadBrandteams           PayloadBrand = "Teams"
	PayloadBrandtesco           PayloadBrand = "Tesco"
	PayloadBrandunknown         PayloadBrand = "Unknown"
	PayloadBrandwellsFargo      PayloadBrand = "WellsFargo"
	PayloadBrandzoom            PayloadBrand = "Zoom"
)

func PossibleValuesForPayloadBrand() []string {
	return []string{
		string(PayloadBrandadobe),
		string(PayloadBrandamericanExpress),
		string(PayloadBrandcapitalOne),
		string(PayloadBranddhl),
		string(PayloadBranddocuSign),
		string(PayloadBranddropbox),
		string(PayloadBrandfacebook),
		string(PayloadBrandfirstAmerican),
		string(PayloadBrandmicrosoft),
		string(PayloadBrandnetflix),
		string(PayloadBrandother),
		string(PayloadBrandscotiabank),
		string(PayloadBrandsendGrid),
		string(PayloadBrandstewartTitle),
		string(PayloadBrandsyrinxCloud),
		string(PayloadBrandteams),
		string(PayloadBrandtesco),
		string(PayloadBrandunknown),
		string(PayloadBrandwellsFargo),
		string(PayloadBrandzoom),
	}
}

func parsePayloadBrand(input string) (*PayloadBrand, error) {
	vals := map[string]PayloadBrand{
		"adobe":           PayloadBrandadobe,
		"americanexpress": PayloadBrandamericanExpress,
		"capitalone":      PayloadBrandcapitalOne,
		"dhl":             PayloadBranddhl,
		"docusign":        PayloadBranddocuSign,
		"dropbox":         PayloadBranddropbox,
		"facebook":        PayloadBrandfacebook,
		"firstamerican":   PayloadBrandfirstAmerican,
		"microsoft":       PayloadBrandmicrosoft,
		"netflix":         PayloadBrandnetflix,
		"other":           PayloadBrandother,
		"scotiabank":      PayloadBrandscotiabank,
		"sendgrid":        PayloadBrandsendGrid,
		"stewarttitle":    PayloadBrandstewartTitle,
		"syrinxcloud":     PayloadBrandsyrinxCloud,
		"teams":           PayloadBrandteams,
		"tesco":           PayloadBrandtesco,
		"unknown":         PayloadBrandunknown,
		"wellsfargo":      PayloadBrandwellsFargo,
		"zoom":            PayloadBrandzoom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadBrand(input)
	return &out, nil
}
