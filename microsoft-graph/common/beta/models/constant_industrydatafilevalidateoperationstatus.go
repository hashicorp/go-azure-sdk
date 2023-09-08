package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFileValidateOperationStatus string

const (
	IndustryDataFileValidateOperationStatusfailed     IndustryDataFileValidateOperationStatus = "Failed"
	IndustryDataFileValidateOperationStatusnotStarted IndustryDataFileValidateOperationStatus = "NotStarted"
	IndustryDataFileValidateOperationStatusrunning    IndustryDataFileValidateOperationStatus = "Running"
	IndustryDataFileValidateOperationStatussucceeded  IndustryDataFileValidateOperationStatus = "Succeeded"
)

func PossibleValuesForIndustryDataFileValidateOperationStatus() []string {
	return []string{
		string(IndustryDataFileValidateOperationStatusfailed),
		string(IndustryDataFileValidateOperationStatusnotStarted),
		string(IndustryDataFileValidateOperationStatusrunning),
		string(IndustryDataFileValidateOperationStatussucceeded),
	}
}

func parseIndustryDataFileValidateOperationStatus(input string) (*IndustryDataFileValidateOperationStatus, error) {
	vals := map[string]IndustryDataFileValidateOperationStatus{
		"failed":     IndustryDataFileValidateOperationStatusfailed,
		"notstarted": IndustryDataFileValidateOperationStatusnotStarted,
		"running":    IndustryDataFileValidateOperationStatusrunning,
		"succeeded":  IndustryDataFileValidateOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataFileValidateOperationStatus(input)
	return &out, nil
}
