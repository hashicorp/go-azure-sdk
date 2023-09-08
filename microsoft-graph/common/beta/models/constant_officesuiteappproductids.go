package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppProductIds string

const (
	OfficeSuiteAppProductIdso365BusinessRetail OfficeSuiteAppProductIds = "O365BusinessRetail"
	OfficeSuiteAppProductIdso365ProPlusRetail  OfficeSuiteAppProductIds = "O365ProPlusRetail"
	OfficeSuiteAppProductIdsprojectProRetail   OfficeSuiteAppProductIds = "ProjectProRetail"
	OfficeSuiteAppProductIdsvisioProRetail     OfficeSuiteAppProductIds = "VisioProRetail"
)

func PossibleValuesForOfficeSuiteAppProductIds() []string {
	return []string{
		string(OfficeSuiteAppProductIdso365BusinessRetail),
		string(OfficeSuiteAppProductIdso365ProPlusRetail),
		string(OfficeSuiteAppProductIdsprojectProRetail),
		string(OfficeSuiteAppProductIdsvisioProRetail),
	}
}

func parseOfficeSuiteAppProductIds(input string) (*OfficeSuiteAppProductIds, error) {
	vals := map[string]OfficeSuiteAppProductIds{
		"o365businessretail": OfficeSuiteAppProductIdso365BusinessRetail,
		"o365proplusretail":  OfficeSuiteAppProductIdso365ProPlusRetail,
		"projectproretail":   OfficeSuiteAppProductIdsprojectProRetail,
		"visioproretail":     OfficeSuiteAppProductIdsvisioProRetail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppProductIds(input)
	return &out, nil
}
