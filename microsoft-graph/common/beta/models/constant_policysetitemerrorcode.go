package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetItemErrorCode string

const (
	PolicySetItemErrorCodedeleted      PolicySetItemErrorCode = "Deleted"
	PolicySetItemErrorCodenoError      PolicySetItemErrorCode = "NoError"
	PolicySetItemErrorCodenotFound     PolicySetItemErrorCode = "NotFound"
	PolicySetItemErrorCodeunauthorized PolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForPolicySetItemErrorCode() []string {
	return []string{
		string(PolicySetItemErrorCodedeleted),
		string(PolicySetItemErrorCodenoError),
		string(PolicySetItemErrorCodenotFound),
		string(PolicySetItemErrorCodeunauthorized),
	}
}

func parsePolicySetItemErrorCode(input string) (*PolicySetItemErrorCode, error) {
	vals := map[string]PolicySetItemErrorCode{
		"deleted":      PolicySetItemErrorCodedeleted,
		"noerror":      PolicySetItemErrorCodenoError,
		"notfound":     PolicySetItemErrorCodenotFound,
		"unauthorized": PolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetItemErrorCode(input)
	return &out, nil
}
