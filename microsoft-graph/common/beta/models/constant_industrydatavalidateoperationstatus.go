package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataValidateOperationStatus string

const (
	IndustryDataValidateOperationStatusfailed     IndustryDataValidateOperationStatus = "Failed"
	IndustryDataValidateOperationStatusnotStarted IndustryDataValidateOperationStatus = "NotStarted"
	IndustryDataValidateOperationStatusrunning    IndustryDataValidateOperationStatus = "Running"
	IndustryDataValidateOperationStatussucceeded  IndustryDataValidateOperationStatus = "Succeeded"
)

func PossibleValuesForIndustryDataValidateOperationStatus() []string {
	return []string{
		string(IndustryDataValidateOperationStatusfailed),
		string(IndustryDataValidateOperationStatusnotStarted),
		string(IndustryDataValidateOperationStatusrunning),
		string(IndustryDataValidateOperationStatussucceeded),
	}
}

func parseIndustryDataValidateOperationStatus(input string) (*IndustryDataValidateOperationStatus, error) {
	vals := map[string]IndustryDataValidateOperationStatus{
		"failed":     IndustryDataValidateOperationStatusfailed,
		"notstarted": IndustryDataValidateOperationStatusnotStarted,
		"running":    IndustryDataValidateOperationStatusrunning,
		"succeeded":  IndustryDataValidateOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataValidateOperationStatus(input)
	return &out, nil
}
