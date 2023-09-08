package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetailsStatus string

const (
	SecurityRunDetailsStatuscompleted       SecurityRunDetailsStatus = "Completed"
	SecurityRunDetailsStatusfailed          SecurityRunDetailsStatus = "Failed"
	SecurityRunDetailsStatuspartiallyFailed SecurityRunDetailsStatus = "PartiallyFailed"
	SecurityRunDetailsStatusrunning         SecurityRunDetailsStatus = "Running"
)

func PossibleValuesForSecurityRunDetailsStatus() []string {
	return []string{
		string(SecurityRunDetailsStatuscompleted),
		string(SecurityRunDetailsStatusfailed),
		string(SecurityRunDetailsStatuspartiallyFailed),
		string(SecurityRunDetailsStatusrunning),
	}
}

func parseSecurityRunDetailsStatus(input string) (*SecurityRunDetailsStatus, error) {
	vals := map[string]SecurityRunDetailsStatus{
		"completed":       SecurityRunDetailsStatuscompleted,
		"failed":          SecurityRunDetailsStatusfailed,
		"partiallyfailed": SecurityRunDetailsStatuspartiallyFailed,
		"running":         SecurityRunDetailsStatusrunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRunDetailsStatus(input)
	return &out, nil
}
