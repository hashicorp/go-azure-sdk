package checkpolicyrestrictions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FieldRestrictionResult string

const (
	FieldRestrictionResultAudit    FieldRestrictionResult = "Audit"
	FieldRestrictionResultDeny     FieldRestrictionResult = "Deny"
	FieldRestrictionResultRemoved  FieldRestrictionResult = "Removed"
	FieldRestrictionResultRequired FieldRestrictionResult = "Required"
)

func PossibleValuesForFieldRestrictionResult() []string {
	return []string{
		string(FieldRestrictionResultAudit),
		string(FieldRestrictionResultDeny),
		string(FieldRestrictionResultRemoved),
		string(FieldRestrictionResultRequired),
	}
}

func parseFieldRestrictionResult(input string) (*FieldRestrictionResult, error) {
	vals := map[string]FieldRestrictionResult{
		"audit":    FieldRestrictionResultAudit,
		"deny":     FieldRestrictionResultDeny,
		"removed":  FieldRestrictionResultRemoved,
		"required": FieldRestrictionResultRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FieldRestrictionResult(input)
	return &out, nil
}
