package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryLegalHoldStatus string

const (
	EdiscoveryLegalHoldStatusError   EdiscoveryLegalHoldStatus = "Error"
	EdiscoveryLegalHoldStatusPending EdiscoveryLegalHoldStatus = "Pending"
	EdiscoveryLegalHoldStatusSuccess EdiscoveryLegalHoldStatus = "Success"
)

func PossibleValuesForEdiscoveryLegalHoldStatus() []string {
	return []string{
		string(EdiscoveryLegalHoldStatusError),
		string(EdiscoveryLegalHoldStatusPending),
		string(EdiscoveryLegalHoldStatusSuccess),
	}
}

func parseEdiscoveryLegalHoldStatus(input string) (*EdiscoveryLegalHoldStatus, error) {
	vals := map[string]EdiscoveryLegalHoldStatus{
		"error":   EdiscoveryLegalHoldStatusError,
		"pending": EdiscoveryLegalHoldStatusPending,
		"success": EdiscoveryLegalHoldStatusSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryLegalHoldStatus(input)
	return &out, nil
}
