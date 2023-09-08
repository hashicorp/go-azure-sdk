package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceContainerStatus string

const (
	EdiscoveryDataSourceContainerStatusActive   EdiscoveryDataSourceContainerStatus = "Active"
	EdiscoveryDataSourceContainerStatusReleased EdiscoveryDataSourceContainerStatus = "Released"
)

func PossibleValuesForEdiscoveryDataSourceContainerStatus() []string {
	return []string{
		string(EdiscoveryDataSourceContainerStatusActive),
		string(EdiscoveryDataSourceContainerStatusReleased),
	}
}

func parseEdiscoveryDataSourceContainerStatus(input string) (*EdiscoveryDataSourceContainerStatus, error) {
	vals := map[string]EdiscoveryDataSourceContainerStatus{
		"active":   EdiscoveryDataSourceContainerStatusActive,
		"released": EdiscoveryDataSourceContainerStatusReleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryDataSourceContainerStatus(input)
	return &out, nil
}
