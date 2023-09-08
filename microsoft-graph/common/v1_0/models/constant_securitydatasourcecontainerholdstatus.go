package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceContainerHoldStatus string

const (
	SecurityDataSourceContainerHoldStatusapplied    SecurityDataSourceContainerHoldStatus = "Applied"
	SecurityDataSourceContainerHoldStatusapplying   SecurityDataSourceContainerHoldStatus = "Applying"
	SecurityDataSourceContainerHoldStatusnotApplied SecurityDataSourceContainerHoldStatus = "NotApplied"
	SecurityDataSourceContainerHoldStatuspartial    SecurityDataSourceContainerHoldStatus = "Partial"
	SecurityDataSourceContainerHoldStatusremoving   SecurityDataSourceContainerHoldStatus = "Removing"
)

func PossibleValuesForSecurityDataSourceContainerHoldStatus() []string {
	return []string{
		string(SecurityDataSourceContainerHoldStatusapplied),
		string(SecurityDataSourceContainerHoldStatusapplying),
		string(SecurityDataSourceContainerHoldStatusnotApplied),
		string(SecurityDataSourceContainerHoldStatuspartial),
		string(SecurityDataSourceContainerHoldStatusremoving),
	}
}

func parseSecurityDataSourceContainerHoldStatus(input string) (*SecurityDataSourceContainerHoldStatus, error) {
	vals := map[string]SecurityDataSourceContainerHoldStatus{
		"applied":    SecurityDataSourceContainerHoldStatusapplied,
		"applying":   SecurityDataSourceContainerHoldStatusapplying,
		"notapplied": SecurityDataSourceContainerHoldStatusnotApplied,
		"partial":    SecurityDataSourceContainerHoldStatuspartial,
		"removing":   SecurityDataSourceContainerHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDataSourceContainerHoldStatus(input)
	return &out, nil
}
