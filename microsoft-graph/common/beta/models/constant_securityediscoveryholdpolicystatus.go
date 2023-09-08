package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryHoldPolicyStatus string

const (
	SecurityEdiscoveryHoldPolicyStatuserror   SecurityEdiscoveryHoldPolicyStatus = "Error"
	SecurityEdiscoveryHoldPolicyStatuspending SecurityEdiscoveryHoldPolicyStatus = "Pending"
	SecurityEdiscoveryHoldPolicyStatussuccess SecurityEdiscoveryHoldPolicyStatus = "Success"
)

func PossibleValuesForSecurityEdiscoveryHoldPolicyStatus() []string {
	return []string{
		string(SecurityEdiscoveryHoldPolicyStatuserror),
		string(SecurityEdiscoveryHoldPolicyStatuspending),
		string(SecurityEdiscoveryHoldPolicyStatussuccess),
	}
}

func parseSecurityEdiscoveryHoldPolicyStatus(input string) (*SecurityEdiscoveryHoldPolicyStatus, error) {
	vals := map[string]SecurityEdiscoveryHoldPolicyStatus{
		"error":   SecurityEdiscoveryHoldPolicyStatuserror,
		"pending": SecurityEdiscoveryHoldPolicyStatuspending,
		"success": SecurityEdiscoveryHoldPolicyStatussuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryHoldPolicyStatus(input)
	return &out, nil
}
