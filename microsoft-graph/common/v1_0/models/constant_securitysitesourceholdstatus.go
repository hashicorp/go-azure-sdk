package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySiteSourceHoldStatus string

const (
	SecuritySiteSourceHoldStatusapplied    SecuritySiteSourceHoldStatus = "Applied"
	SecuritySiteSourceHoldStatusapplying   SecuritySiteSourceHoldStatus = "Applying"
	SecuritySiteSourceHoldStatusnotApplied SecuritySiteSourceHoldStatus = "NotApplied"
	SecuritySiteSourceHoldStatuspartial    SecuritySiteSourceHoldStatus = "Partial"
	SecuritySiteSourceHoldStatusremoving   SecuritySiteSourceHoldStatus = "Removing"
)

func PossibleValuesForSecuritySiteSourceHoldStatus() []string {
	return []string{
		string(SecuritySiteSourceHoldStatusapplied),
		string(SecuritySiteSourceHoldStatusapplying),
		string(SecuritySiteSourceHoldStatusnotApplied),
		string(SecuritySiteSourceHoldStatuspartial),
		string(SecuritySiteSourceHoldStatusremoving),
	}
}

func parseSecuritySiteSourceHoldStatus(input string) (*SecuritySiteSourceHoldStatus, error) {
	vals := map[string]SecuritySiteSourceHoldStatus{
		"applied":    SecuritySiteSourceHoldStatusapplied,
		"applying":   SecuritySiteSourceHoldStatusapplying,
		"notapplied": SecuritySiteSourceHoldStatusnotApplied,
		"partial":    SecuritySiteSourceHoldStatuspartial,
		"removing":   SecuritySiteSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySiteSourceHoldStatus(input)
	return &out, nil
}
