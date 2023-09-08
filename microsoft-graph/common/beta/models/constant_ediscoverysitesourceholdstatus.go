package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySiteSourceHoldStatus string

const (
	EdiscoverySiteSourceHoldStatusapplied    EdiscoverySiteSourceHoldStatus = "Applied"
	EdiscoverySiteSourceHoldStatusapplying   EdiscoverySiteSourceHoldStatus = "Applying"
	EdiscoverySiteSourceHoldStatusnotApplied EdiscoverySiteSourceHoldStatus = "NotApplied"
	EdiscoverySiteSourceHoldStatuspartial    EdiscoverySiteSourceHoldStatus = "Partial"
	EdiscoverySiteSourceHoldStatusremoving   EdiscoverySiteSourceHoldStatus = "Removing"
)

func PossibleValuesForEdiscoverySiteSourceHoldStatus() []string {
	return []string{
		string(EdiscoverySiteSourceHoldStatusapplied),
		string(EdiscoverySiteSourceHoldStatusapplying),
		string(EdiscoverySiteSourceHoldStatusnotApplied),
		string(EdiscoverySiteSourceHoldStatuspartial),
		string(EdiscoverySiteSourceHoldStatusremoving),
	}
}

func parseEdiscoverySiteSourceHoldStatus(input string) (*EdiscoverySiteSourceHoldStatus, error) {
	vals := map[string]EdiscoverySiteSourceHoldStatus{
		"applied":    EdiscoverySiteSourceHoldStatusapplied,
		"applying":   EdiscoverySiteSourceHoldStatusapplying,
		"notapplied": EdiscoverySiteSourceHoldStatusnotApplied,
		"partial":    EdiscoverySiteSourceHoldStatuspartial,
		"removing":   EdiscoverySiteSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoverySiteSourceHoldStatus(input)
	return &out, nil
}
