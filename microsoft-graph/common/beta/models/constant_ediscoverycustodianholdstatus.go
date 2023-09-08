package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCustodianHoldStatus string

const (
	EdiscoveryCustodianHoldStatusapplied    EdiscoveryCustodianHoldStatus = "Applied"
	EdiscoveryCustodianHoldStatusapplying   EdiscoveryCustodianHoldStatus = "Applying"
	EdiscoveryCustodianHoldStatusnotApplied EdiscoveryCustodianHoldStatus = "NotApplied"
	EdiscoveryCustodianHoldStatuspartial    EdiscoveryCustodianHoldStatus = "Partial"
	EdiscoveryCustodianHoldStatusremoving   EdiscoveryCustodianHoldStatus = "Removing"
)

func PossibleValuesForEdiscoveryCustodianHoldStatus() []string {
	return []string{
		string(EdiscoveryCustodianHoldStatusapplied),
		string(EdiscoveryCustodianHoldStatusapplying),
		string(EdiscoveryCustodianHoldStatusnotApplied),
		string(EdiscoveryCustodianHoldStatuspartial),
		string(EdiscoveryCustodianHoldStatusremoving),
	}
}

func parseEdiscoveryCustodianHoldStatus(input string) (*EdiscoveryCustodianHoldStatus, error) {
	vals := map[string]EdiscoveryCustodianHoldStatus{
		"applied":    EdiscoveryCustodianHoldStatusapplied,
		"applying":   EdiscoveryCustodianHoldStatusapplying,
		"notapplied": EdiscoveryCustodianHoldStatusnotApplied,
		"partial":    EdiscoveryCustodianHoldStatuspartial,
		"removing":   EdiscoveryCustodianHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCustodianHoldStatus(input)
	return &out, nil
}
