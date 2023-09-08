package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryNoncustodialDataSourceHoldStatus string

const (
	SecurityEdiscoveryNoncustodialDataSourceHoldStatusapplied    SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "Applied"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatusapplying   SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "Applying"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatusnotApplied SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "NotApplied"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatuspartial    SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "Partial"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatusremoving   SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "Removing"
)

func PossibleValuesForSecurityEdiscoveryNoncustodialDataSourceHoldStatus() []string {
	return []string{
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatusapplied),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatusapplying),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatusnotApplied),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatuspartial),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatusremoving),
	}
}

func parseSecurityEdiscoveryNoncustodialDataSourceHoldStatus(input string) (*SecurityEdiscoveryNoncustodialDataSourceHoldStatus, error) {
	vals := map[string]SecurityEdiscoveryNoncustodialDataSourceHoldStatus{
		"applied":    SecurityEdiscoveryNoncustodialDataSourceHoldStatusapplied,
		"applying":   SecurityEdiscoveryNoncustodialDataSourceHoldStatusapplying,
		"notapplied": SecurityEdiscoveryNoncustodialDataSourceHoldStatusnotApplied,
		"partial":    SecurityEdiscoveryNoncustodialDataSourceHoldStatuspartial,
		"removing":   SecurityEdiscoveryNoncustodialDataSourceHoldStatusremoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryNoncustodialDataSourceHoldStatus(input)
	return &out, nil
}
