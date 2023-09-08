package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUnifiedGroupSourceHoldStatus string

const (
	SecurityUnifiedGroupSourceHoldStatusapplied    SecurityUnifiedGroupSourceHoldStatus = "Applied"
	SecurityUnifiedGroupSourceHoldStatusapplying   SecurityUnifiedGroupSourceHoldStatus = "Applying"
	SecurityUnifiedGroupSourceHoldStatusnotApplied SecurityUnifiedGroupSourceHoldStatus = "NotApplied"
	SecurityUnifiedGroupSourceHoldStatuspartial    SecurityUnifiedGroupSourceHoldStatus = "Partial"
	SecurityUnifiedGroupSourceHoldStatusremoving   SecurityUnifiedGroupSourceHoldStatus = "Removing"
)

func PossibleValuesForSecurityUnifiedGroupSourceHoldStatus() []string {
	return []string{
		string(SecurityUnifiedGroupSourceHoldStatusapplied),
		string(SecurityUnifiedGroupSourceHoldStatusapplying),
		string(SecurityUnifiedGroupSourceHoldStatusnotApplied),
		string(SecurityUnifiedGroupSourceHoldStatuspartial),
		string(SecurityUnifiedGroupSourceHoldStatusremoving),
	}
}

func parseSecurityUnifiedGroupSourceHoldStatus(input string) (*SecurityUnifiedGroupSourceHoldStatus, error) {
	vals := map[string]SecurityUnifiedGroupSourceHoldStatus{
		"applied":    SecurityUnifiedGroupSourceHoldStatusapplied,
		"applying":   SecurityUnifiedGroupSourceHoldStatusapplying,
		"notapplied": SecurityUnifiedGroupSourceHoldStatusnotApplied,
		"partial":    SecurityUnifiedGroupSourceHoldStatuspartial,
		"removing":   SecurityUnifiedGroupSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUnifiedGroupSourceHoldStatus(input)
	return &out, nil
}
