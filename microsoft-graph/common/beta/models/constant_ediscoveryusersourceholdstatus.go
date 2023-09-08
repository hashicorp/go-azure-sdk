package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUserSourceHoldStatus string

const (
	EdiscoveryUserSourceHoldStatusapplied    EdiscoveryUserSourceHoldStatus = "Applied"
	EdiscoveryUserSourceHoldStatusapplying   EdiscoveryUserSourceHoldStatus = "Applying"
	EdiscoveryUserSourceHoldStatusnotApplied EdiscoveryUserSourceHoldStatus = "NotApplied"
	EdiscoveryUserSourceHoldStatuspartial    EdiscoveryUserSourceHoldStatus = "Partial"
	EdiscoveryUserSourceHoldStatusremoving   EdiscoveryUserSourceHoldStatus = "Removing"
)

func PossibleValuesForEdiscoveryUserSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryUserSourceHoldStatusapplied),
		string(EdiscoveryUserSourceHoldStatusapplying),
		string(EdiscoveryUserSourceHoldStatusnotApplied),
		string(EdiscoveryUserSourceHoldStatuspartial),
		string(EdiscoveryUserSourceHoldStatusremoving),
	}
}

func parseEdiscoveryUserSourceHoldStatus(input string) (*EdiscoveryUserSourceHoldStatus, error) {
	vals := map[string]EdiscoveryUserSourceHoldStatus{
		"applied":    EdiscoveryUserSourceHoldStatusapplied,
		"applying":   EdiscoveryUserSourceHoldStatusapplying,
		"notapplied": EdiscoveryUserSourceHoldStatusnotApplied,
		"partial":    EdiscoveryUserSourceHoldStatuspartial,
		"removing":   EdiscoveryUserSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUserSourceHoldStatus(input)
	return &out, nil
}
