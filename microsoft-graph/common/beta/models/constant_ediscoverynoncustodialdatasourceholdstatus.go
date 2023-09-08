package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryNoncustodialDataSourceHoldStatus string

const (
	EdiscoveryNoncustodialDataSourceHoldStatusapplied    EdiscoveryNoncustodialDataSourceHoldStatus = "Applied"
	EdiscoveryNoncustodialDataSourceHoldStatusapplying   EdiscoveryNoncustodialDataSourceHoldStatus = "Applying"
	EdiscoveryNoncustodialDataSourceHoldStatusnotApplied EdiscoveryNoncustodialDataSourceHoldStatus = "NotApplied"
	EdiscoveryNoncustodialDataSourceHoldStatuspartial    EdiscoveryNoncustodialDataSourceHoldStatus = "Partial"
	EdiscoveryNoncustodialDataSourceHoldStatusremoving   EdiscoveryNoncustodialDataSourceHoldStatus = "Removing"
)

func PossibleValuesForEdiscoveryNoncustodialDataSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryNoncustodialDataSourceHoldStatusapplied),
		string(EdiscoveryNoncustodialDataSourceHoldStatusapplying),
		string(EdiscoveryNoncustodialDataSourceHoldStatusnotApplied),
		string(EdiscoveryNoncustodialDataSourceHoldStatuspartial),
		string(EdiscoveryNoncustodialDataSourceHoldStatusremoving),
	}
}

func parseEdiscoveryNoncustodialDataSourceHoldStatus(input string) (*EdiscoveryNoncustodialDataSourceHoldStatus, error) {
	vals := map[string]EdiscoveryNoncustodialDataSourceHoldStatus{
		"applied":    EdiscoveryNoncustodialDataSourceHoldStatusapplied,
		"applying":   EdiscoveryNoncustodialDataSourceHoldStatusapplying,
		"notapplied": EdiscoveryNoncustodialDataSourceHoldStatusnotApplied,
		"partial":    EdiscoveryNoncustodialDataSourceHoldStatuspartial,
		"removing":   EdiscoveryNoncustodialDataSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryNoncustodialDataSourceHoldStatus(input)
	return &out, nil
}
