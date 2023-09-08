package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceContainerStatus string

const (
	SecurityDataSourceContainerStatusactive   SecurityDataSourceContainerStatus = "Active"
	SecurityDataSourceContainerStatusreleased SecurityDataSourceContainerStatus = "Released"
)

func PossibleValuesForSecurityDataSourceContainerStatus() []string {
	return []string{
		string(SecurityDataSourceContainerStatusactive),
		string(SecurityDataSourceContainerStatusreleased),
	}
}

func parseSecurityDataSourceContainerStatus(input string) (*SecurityDataSourceContainerStatus, error) {
	vals := map[string]SecurityDataSourceContainerStatus{
		"active":   SecurityDataSourceContainerStatusactive,
		"released": SecurityDataSourceContainerStatusreleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDataSourceContainerStatus(input)
	return &out, nil
}
