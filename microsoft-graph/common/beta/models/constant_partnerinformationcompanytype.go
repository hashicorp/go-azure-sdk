package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerInformationCompanyType string

const (
	PartnerInformationCompanyTypebreadthPartner                          PartnerInformationCompanyType = "BreadthPartner"
	PartnerInformationCompanyTypebreadthPartnerDelegatedAdmin            PartnerInformationCompanyType = "BreadthPartnerDelegatedAdmin"
	PartnerInformationCompanyTypemicrosoftSupport                        PartnerInformationCompanyType = "MicrosoftSupport"
	PartnerInformationCompanyTyperesellerPartnerDelegatedAdmin           PartnerInformationCompanyType = "ResellerPartnerDelegatedAdmin"
	PartnerInformationCompanyTypesyndicatePartner                        PartnerInformationCompanyType = "SyndicatePartner"
	PartnerInformationCompanyTypevalueAddedResellerPartnerDelegatedAdmin PartnerInformationCompanyType = "ValueAddedResellerPartnerDelegatedAdmin"
)

func PossibleValuesForPartnerInformationCompanyType() []string {
	return []string{
		string(PartnerInformationCompanyTypebreadthPartner),
		string(PartnerInformationCompanyTypebreadthPartnerDelegatedAdmin),
		string(PartnerInformationCompanyTypemicrosoftSupport),
		string(PartnerInformationCompanyTyperesellerPartnerDelegatedAdmin),
		string(PartnerInformationCompanyTypesyndicatePartner),
		string(PartnerInformationCompanyTypevalueAddedResellerPartnerDelegatedAdmin),
	}
}

func parsePartnerInformationCompanyType(input string) (*PartnerInformationCompanyType, error) {
	vals := map[string]PartnerInformationCompanyType{
		"breadthpartner":                          PartnerInformationCompanyTypebreadthPartner,
		"breadthpartnerdelegatedadmin":            PartnerInformationCompanyTypebreadthPartnerDelegatedAdmin,
		"microsoftsupport":                        PartnerInformationCompanyTypemicrosoftSupport,
		"resellerpartnerdelegatedadmin":           PartnerInformationCompanyTyperesellerPartnerDelegatedAdmin,
		"syndicatepartner":                        PartnerInformationCompanyTypesyndicatePartner,
		"valueaddedresellerpartnerdelegatedadmin": PartnerInformationCompanyTypevalueAddedResellerPartnerDelegatedAdmin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerInformationCompanyType(input)
	return &out, nil
}
