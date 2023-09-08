package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPolicyBaseStatus string

const (
	SecurityPolicyBaseStatuserror   SecurityPolicyBaseStatus = "Error"
	SecurityPolicyBaseStatuspending SecurityPolicyBaseStatus = "Pending"
	SecurityPolicyBaseStatussuccess SecurityPolicyBaseStatus = "Success"
)

func PossibleValuesForSecurityPolicyBaseStatus() []string {
	return []string{
		string(SecurityPolicyBaseStatuserror),
		string(SecurityPolicyBaseStatuspending),
		string(SecurityPolicyBaseStatussuccess),
	}
}

func parseSecurityPolicyBaseStatus(input string) (*SecurityPolicyBaseStatus, error) {
	vals := map[string]SecurityPolicyBaseStatus{
		"error":   SecurityPolicyBaseStatuserror,
		"pending": SecurityPolicyBaseStatuspending,
		"success": SecurityPolicyBaseStatussuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityPolicyBaseStatus(input)
	return &out, nil
}
