package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUnifiedGroupSourceHoldStatus string

const (
	EdiscoveryUnifiedGroupSourceHoldStatusapplied    EdiscoveryUnifiedGroupSourceHoldStatus = "Applied"
	EdiscoveryUnifiedGroupSourceHoldStatusapplying   EdiscoveryUnifiedGroupSourceHoldStatus = "Applying"
	EdiscoveryUnifiedGroupSourceHoldStatusnotApplied EdiscoveryUnifiedGroupSourceHoldStatus = "NotApplied"
	EdiscoveryUnifiedGroupSourceHoldStatuspartial    EdiscoveryUnifiedGroupSourceHoldStatus = "Partial"
	EdiscoveryUnifiedGroupSourceHoldStatusremoving   EdiscoveryUnifiedGroupSourceHoldStatus = "Removing"
)

func PossibleValuesForEdiscoveryUnifiedGroupSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryUnifiedGroupSourceHoldStatusapplied),
		string(EdiscoveryUnifiedGroupSourceHoldStatusapplying),
		string(EdiscoveryUnifiedGroupSourceHoldStatusnotApplied),
		string(EdiscoveryUnifiedGroupSourceHoldStatuspartial),
		string(EdiscoveryUnifiedGroupSourceHoldStatusremoving),
	}
}

func parseEdiscoveryUnifiedGroupSourceHoldStatus(input string) (*EdiscoveryUnifiedGroupSourceHoldStatus, error) {
	vals := map[string]EdiscoveryUnifiedGroupSourceHoldStatus{
		"applied":    EdiscoveryUnifiedGroupSourceHoldStatusapplied,
		"applying":   EdiscoveryUnifiedGroupSourceHoldStatusapplying,
		"notapplied": EdiscoveryUnifiedGroupSourceHoldStatusnotApplied,
		"partial":    EdiscoveryUnifiedGroupSourceHoldStatuspartial,
		"removing":   EdiscoveryUnifiedGroupSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUnifiedGroupSourceHoldStatus(input)
	return &out, nil
}
