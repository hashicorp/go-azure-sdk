package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCustodianHoldStatus string

const (
	SecurityEdiscoveryCustodianHoldStatusapplied    SecurityEdiscoveryCustodianHoldStatus = "Applied"
	SecurityEdiscoveryCustodianHoldStatusapplying   SecurityEdiscoveryCustodianHoldStatus = "Applying"
	SecurityEdiscoveryCustodianHoldStatusnotApplied SecurityEdiscoveryCustodianHoldStatus = "NotApplied"
	SecurityEdiscoveryCustodianHoldStatuspartial    SecurityEdiscoveryCustodianHoldStatus = "Partial"
	SecurityEdiscoveryCustodianHoldStatusremoving   SecurityEdiscoveryCustodianHoldStatus = "Removing"
)

func PossibleValuesForSecurityEdiscoveryCustodianHoldStatus() []string {
	return []string{
		string(SecurityEdiscoveryCustodianHoldStatusapplied),
		string(SecurityEdiscoveryCustodianHoldStatusapplying),
		string(SecurityEdiscoveryCustodianHoldStatusnotApplied),
		string(SecurityEdiscoveryCustodianHoldStatuspartial),
		string(SecurityEdiscoveryCustodianHoldStatusremoving),
	}
}

func parseSecurityEdiscoveryCustodianHoldStatus(input string) (*SecurityEdiscoveryCustodianHoldStatus, error) {
	vals := map[string]SecurityEdiscoveryCustodianHoldStatus{
		"applied":    SecurityEdiscoveryCustodianHoldStatusapplied,
		"applying":   SecurityEdiscoveryCustodianHoldStatusapplying,
		"notapplied": SecurityEdiscoveryCustodianHoldStatusnotApplied,
		"partial":    SecurityEdiscoveryCustodianHoldStatuspartial,
		"removing":   SecurityEdiscoveryCustodianHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryCustodianHoldStatus(input)
	return &out, nil
}
