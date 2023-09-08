package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetStatus string

const (
	PolicySetStatuserror          PolicySetStatus = "Error"
	PolicySetStatusnotAssigned    PolicySetStatus = "NotAssigned"
	PolicySetStatuspartialSuccess PolicySetStatus = "PartialSuccess"
	PolicySetStatussuccess        PolicySetStatus = "Success"
	PolicySetStatusunknown        PolicySetStatus = "Unknown"
	PolicySetStatusvalidating     PolicySetStatus = "Validating"
)

func PossibleValuesForPolicySetStatus() []string {
	return []string{
		string(PolicySetStatuserror),
		string(PolicySetStatusnotAssigned),
		string(PolicySetStatuspartialSuccess),
		string(PolicySetStatussuccess),
		string(PolicySetStatusunknown),
		string(PolicySetStatusvalidating),
	}
}

func parsePolicySetStatus(input string) (*PolicySetStatus, error) {
	vals := map[string]PolicySetStatus{
		"error":          PolicySetStatuserror,
		"notassigned":    PolicySetStatusnotAssigned,
		"partialsuccess": PolicySetStatuspartialSuccess,
		"success":        PolicySetStatussuccess,
		"unknown":        PolicySetStatusunknown,
		"validating":     PolicySetStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetStatus(input)
	return &out, nil
}
