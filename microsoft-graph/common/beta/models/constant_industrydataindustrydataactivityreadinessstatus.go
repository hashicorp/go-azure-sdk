package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityReadinessStatus string

const (
	IndustryDataIndustryDataActivityReadinessStatusdisabled IndustryDataIndustryDataActivityReadinessStatus = "Disabled"
	IndustryDataIndustryDataActivityReadinessStatusexpired  IndustryDataIndustryDataActivityReadinessStatus = "Expired"
	IndustryDataIndustryDataActivityReadinessStatusfailed   IndustryDataIndustryDataActivityReadinessStatus = "Failed"
	IndustryDataIndustryDataActivityReadinessStatusnotReady IndustryDataIndustryDataActivityReadinessStatus = "NotReady"
	IndustryDataIndustryDataActivityReadinessStatusready    IndustryDataIndustryDataActivityReadinessStatus = "Ready"
)

func PossibleValuesForIndustryDataIndustryDataActivityReadinessStatus() []string {
	return []string{
		string(IndustryDataIndustryDataActivityReadinessStatusdisabled),
		string(IndustryDataIndustryDataActivityReadinessStatusexpired),
		string(IndustryDataIndustryDataActivityReadinessStatusfailed),
		string(IndustryDataIndustryDataActivityReadinessStatusnotReady),
		string(IndustryDataIndustryDataActivityReadinessStatusready),
	}
}

func parseIndustryDataIndustryDataActivityReadinessStatus(input string) (*IndustryDataIndustryDataActivityReadinessStatus, error) {
	vals := map[string]IndustryDataIndustryDataActivityReadinessStatus{
		"disabled": IndustryDataIndustryDataActivityReadinessStatusdisabled,
		"expired":  IndustryDataIndustryDataActivityReadinessStatusexpired,
		"failed":   IndustryDataIndustryDataActivityReadinessStatusfailed,
		"notready": IndustryDataIndustryDataActivityReadinessStatusnotReady,
		"ready":    IndustryDataIndustryDataActivityReadinessStatusready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataActivityReadinessStatus(input)
	return &out, nil
}
