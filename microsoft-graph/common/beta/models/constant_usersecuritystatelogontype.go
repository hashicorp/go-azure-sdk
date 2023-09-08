package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateLogonType string

const (
	UserSecurityStateLogonTypebatch             UserSecurityStateLogonType = "Batch"
	UserSecurityStateLogonTypeinteractive       UserSecurityStateLogonType = "Interactive"
	UserSecurityStateLogonTypenetwork           UserSecurityStateLogonType = "Network"
	UserSecurityStateLogonTyperemoteInteractive UserSecurityStateLogonType = "RemoteInteractive"
	UserSecurityStateLogonTypeservice           UserSecurityStateLogonType = "Service"
	UserSecurityStateLogonTypeunknown           UserSecurityStateLogonType = "Unknown"
)

func PossibleValuesForUserSecurityStateLogonType() []string {
	return []string{
		string(UserSecurityStateLogonTypebatch),
		string(UserSecurityStateLogonTypeinteractive),
		string(UserSecurityStateLogonTypenetwork),
		string(UserSecurityStateLogonTyperemoteInteractive),
		string(UserSecurityStateLogonTypeservice),
		string(UserSecurityStateLogonTypeunknown),
	}
}

func parseUserSecurityStateLogonType(input string) (*UserSecurityStateLogonType, error) {
	vals := map[string]UserSecurityStateLogonType{
		"batch":             UserSecurityStateLogonTypebatch,
		"interactive":       UserSecurityStateLogonTypeinteractive,
		"network":           UserSecurityStateLogonTypenetwork,
		"remoteinteractive": UserSecurityStateLogonTyperemoteInteractive,
		"service":           UserSecurityStateLogonTypeservice,
		"unknown":           UserSecurityStateLogonTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSecurityStateLogonType(input)
	return &out, nil
}
