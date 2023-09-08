package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPolicySetItemErrorCode string

const (
	MobileAppPolicySetItemErrorCodedeleted      MobileAppPolicySetItemErrorCode = "Deleted"
	MobileAppPolicySetItemErrorCodenoError      MobileAppPolicySetItemErrorCode = "NoError"
	MobileAppPolicySetItemErrorCodenotFound     MobileAppPolicySetItemErrorCode = "NotFound"
	MobileAppPolicySetItemErrorCodeunauthorized MobileAppPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForMobileAppPolicySetItemErrorCode() []string {
	return []string{
		string(MobileAppPolicySetItemErrorCodedeleted),
		string(MobileAppPolicySetItemErrorCodenoError),
		string(MobileAppPolicySetItemErrorCodenotFound),
		string(MobileAppPolicySetItemErrorCodeunauthorized),
	}
}

func parseMobileAppPolicySetItemErrorCode(input string) (*MobileAppPolicySetItemErrorCode, error) {
	vals := map[string]MobileAppPolicySetItemErrorCode{
		"deleted":      MobileAppPolicySetItemErrorCodedeleted,
		"noerror":      MobileAppPolicySetItemErrorCodenoError,
		"notfound":     MobileAppPolicySetItemErrorCodenotFound,
		"unauthorized": MobileAppPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPolicySetItemErrorCode(input)
	return &out, nil
}
