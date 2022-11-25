package customlocations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostType string

const (
	HostTypeKubernetes HostType = "Kubernetes"
)

func PossibleValuesForHostType() []string {
	return []string{
		string(HostTypeKubernetes),
	}
}

func parseHostType(input string) (*HostType, error) {
	vals := map[string]HostType{
		"kubernetes": HostTypeKubernetes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostType(input)
	return &out, nil
}
