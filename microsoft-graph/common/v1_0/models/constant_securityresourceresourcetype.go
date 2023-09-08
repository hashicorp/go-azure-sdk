package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResourceResourceType string

const (
	SecurityResourceResourceTypeattacked SecurityResourceResourceType = "Attacked"
	SecurityResourceResourceTyperelated  SecurityResourceResourceType = "Related"
	SecurityResourceResourceTypeunknown  SecurityResourceResourceType = "Unknown"
)

func PossibleValuesForSecurityResourceResourceType() []string {
	return []string{
		string(SecurityResourceResourceTypeattacked),
		string(SecurityResourceResourceTyperelated),
		string(SecurityResourceResourceTypeunknown),
	}
}

func parseSecurityResourceResourceType(input string) (*SecurityResourceResourceType, error) {
	vals := map[string]SecurityResourceResourceType{
		"attacked": SecurityResourceResourceTypeattacked,
		"related":  SecurityResourceResourceTyperelated,
		"unknown":  SecurityResourceResourceTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityResourceResourceType(input)
	return &out, nil
}
