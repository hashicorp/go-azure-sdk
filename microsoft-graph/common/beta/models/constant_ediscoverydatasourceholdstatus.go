package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceHoldStatus string

const (
	EdiscoveryDataSourceHoldStatusapplied    EdiscoveryDataSourceHoldStatus = "Applied"
	EdiscoveryDataSourceHoldStatusapplying   EdiscoveryDataSourceHoldStatus = "Applying"
	EdiscoveryDataSourceHoldStatusnotApplied EdiscoveryDataSourceHoldStatus = "NotApplied"
	EdiscoveryDataSourceHoldStatuspartial    EdiscoveryDataSourceHoldStatus = "Partial"
	EdiscoveryDataSourceHoldStatusremoving   EdiscoveryDataSourceHoldStatus = "Removing"
)

func PossibleValuesForEdiscoveryDataSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryDataSourceHoldStatusapplied),
		string(EdiscoveryDataSourceHoldStatusapplying),
		string(EdiscoveryDataSourceHoldStatusnotApplied),
		string(EdiscoveryDataSourceHoldStatuspartial),
		string(EdiscoveryDataSourceHoldStatusremoving),
	}
}

func parseEdiscoveryDataSourceHoldStatus(input string) (*EdiscoveryDataSourceHoldStatus, error) {
	vals := map[string]EdiscoveryDataSourceHoldStatus{
		"applied":    EdiscoveryDataSourceHoldStatusapplied,
		"applying":   EdiscoveryDataSourceHoldStatusapplying,
		"notapplied": EdiscoveryDataSourceHoldStatusnotApplied,
		"partial":    EdiscoveryDataSourceHoldStatuspartial,
		"removing":   EdiscoveryDataSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryDataSourceHoldStatus(input)
	return &out, nil
}
