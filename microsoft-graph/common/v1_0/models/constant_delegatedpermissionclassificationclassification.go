package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedPermissionClassificationClassification string

const (
	DelegatedPermissionClassificationClassificationhigh   DelegatedPermissionClassificationClassification = "High"
	DelegatedPermissionClassificationClassificationlow    DelegatedPermissionClassificationClassification = "Low"
	DelegatedPermissionClassificationClassificationmedium DelegatedPermissionClassificationClassification = "Medium"
)

func PossibleValuesForDelegatedPermissionClassificationClassification() []string {
	return []string{
		string(DelegatedPermissionClassificationClassificationhigh),
		string(DelegatedPermissionClassificationClassificationlow),
		string(DelegatedPermissionClassificationClassificationmedium),
	}
}

func parseDelegatedPermissionClassificationClassification(input string) (*DelegatedPermissionClassificationClassification, error) {
	vals := map[string]DelegatedPermissionClassificationClassification{
		"high":   DelegatedPermissionClassificationClassificationhigh,
		"low":    DelegatedPermissionClassificationClassificationlow,
		"medium": DelegatedPermissionClassificationClassificationmedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedPermissionClassificationClassification(input)
	return &out, nil
}
