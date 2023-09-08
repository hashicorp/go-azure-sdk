package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPolicySetItemStatus string

const (
	MobileAppPolicySetItemStatuserror          MobileAppPolicySetItemStatus = "Error"
	MobileAppPolicySetItemStatusnotAssigned    MobileAppPolicySetItemStatus = "NotAssigned"
	MobileAppPolicySetItemStatuspartialSuccess MobileAppPolicySetItemStatus = "PartialSuccess"
	MobileAppPolicySetItemStatussuccess        MobileAppPolicySetItemStatus = "Success"
	MobileAppPolicySetItemStatusunknown        MobileAppPolicySetItemStatus = "Unknown"
	MobileAppPolicySetItemStatusvalidating     MobileAppPolicySetItemStatus = "Validating"
)

func PossibleValuesForMobileAppPolicySetItemStatus() []string {
	return []string{
		string(MobileAppPolicySetItemStatuserror),
		string(MobileAppPolicySetItemStatusnotAssigned),
		string(MobileAppPolicySetItemStatuspartialSuccess),
		string(MobileAppPolicySetItemStatussuccess),
		string(MobileAppPolicySetItemStatusunknown),
		string(MobileAppPolicySetItemStatusvalidating),
	}
}

func parseMobileAppPolicySetItemStatus(input string) (*MobileAppPolicySetItemStatus, error) {
	vals := map[string]MobileAppPolicySetItemStatus{
		"error":          MobileAppPolicySetItemStatuserror,
		"notassigned":    MobileAppPolicySetItemStatusnotAssigned,
		"partialsuccess": MobileAppPolicySetItemStatuspartialSuccess,
		"success":        MobileAppPolicySetItemStatussuccess,
		"unknown":        MobileAppPolicySetItemStatusunknown,
		"validating":     MobileAppPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPolicySetItemStatus(input)
	return &out, nil
}
