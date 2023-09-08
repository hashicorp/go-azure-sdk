package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCustodianStatus string

const (
	SecurityEdiscoveryCustodianStatusactive   SecurityEdiscoveryCustodianStatus = "Active"
	SecurityEdiscoveryCustodianStatusreleased SecurityEdiscoveryCustodianStatus = "Released"
)

func PossibleValuesForSecurityEdiscoveryCustodianStatus() []string {
	return []string{
		string(SecurityEdiscoveryCustodianStatusactive),
		string(SecurityEdiscoveryCustodianStatusreleased),
	}
}

func parseSecurityEdiscoveryCustodianStatus(input string) (*SecurityEdiscoveryCustodianStatus, error) {
	vals := map[string]SecurityEdiscoveryCustodianStatus{
		"active":   SecurityEdiscoveryCustodianStatusactive,
		"released": SecurityEdiscoveryCustodianStatusreleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryCustodianStatus(input)
	return &out, nil
}
