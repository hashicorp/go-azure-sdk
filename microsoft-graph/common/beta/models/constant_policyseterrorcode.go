package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetErrorCode string

const (
	PolicySetErrorCodedeleted      PolicySetErrorCode = "Deleted"
	PolicySetErrorCodenoError      PolicySetErrorCode = "NoError"
	PolicySetErrorCodenotFound     PolicySetErrorCode = "NotFound"
	PolicySetErrorCodeunauthorized PolicySetErrorCode = "Unauthorized"
)

func PossibleValuesForPolicySetErrorCode() []string {
	return []string{
		string(PolicySetErrorCodedeleted),
		string(PolicySetErrorCodenoError),
		string(PolicySetErrorCodenotFound),
		string(PolicySetErrorCodeunauthorized),
	}
}

func parsePolicySetErrorCode(input string) (*PolicySetErrorCode, error) {
	vals := map[string]PolicySetErrorCode{
		"deleted":      PolicySetErrorCodedeleted,
		"noerror":      PolicySetErrorCodenoError,
		"notfound":     PolicySetErrorCodenotFound,
		"unauthorized": PolicySetErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetErrorCode(input)
	return &out, nil
}
