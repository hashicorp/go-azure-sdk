package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleSummaryStatus string

const (
	PrivilegedRoleSummaryStatusbad PrivilegedRoleSummaryStatus = "Bad"
	PrivilegedRoleSummaryStatusok  PrivilegedRoleSummaryStatus = "Ok"
)

func PossibleValuesForPrivilegedRoleSummaryStatus() []string {
	return []string{
		string(PrivilegedRoleSummaryStatusbad),
		string(PrivilegedRoleSummaryStatusok),
	}
}

func parsePrivilegedRoleSummaryStatus(input string) (*PrivilegedRoleSummaryStatus, error) {
	vals := map[string]PrivilegedRoleSummaryStatus{
		"bad": PrivilegedRoleSummaryStatusbad,
		"ok":  PrivilegedRoleSummaryStatusok,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedRoleSummaryStatus(input)
	return &out, nil
}
