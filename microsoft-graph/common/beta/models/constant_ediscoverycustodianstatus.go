package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCustodianStatus string

const (
	EdiscoveryCustodianStatusActive   EdiscoveryCustodianStatus = "Active"
	EdiscoveryCustodianStatusReleased EdiscoveryCustodianStatus = "Released"
)

func PossibleValuesForEdiscoveryCustodianStatus() []string {
	return []string{
		string(EdiscoveryCustodianStatusActive),
		string(EdiscoveryCustodianStatusReleased),
	}
}

func parseEdiscoveryCustodianStatus(input string) (*EdiscoveryCustodianStatus, error) {
	vals := map[string]EdiscoveryCustodianStatus{
		"active":   EdiscoveryCustodianStatusActive,
		"released": EdiscoveryCustodianStatusReleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCustodianStatus(input)
	return &out, nil
}
