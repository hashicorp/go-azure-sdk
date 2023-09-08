package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryIngestionSource string

const (
	GroupPolicyCategoryIngestionSourcebuiltIn GroupPolicyCategoryIngestionSource = "BuiltIn"
	GroupPolicyCategoryIngestionSourcecustom  GroupPolicyCategoryIngestionSource = "Custom"
	GroupPolicyCategoryIngestionSourceunknown GroupPolicyCategoryIngestionSource = "Unknown"
)

func PossibleValuesForGroupPolicyCategoryIngestionSource() []string {
	return []string{
		string(GroupPolicyCategoryIngestionSourcebuiltIn),
		string(GroupPolicyCategoryIngestionSourcecustom),
		string(GroupPolicyCategoryIngestionSourceunknown),
	}
}

func parseGroupPolicyCategoryIngestionSource(input string) (*GroupPolicyCategoryIngestionSource, error) {
	vals := map[string]GroupPolicyCategoryIngestionSource{
		"builtin": GroupPolicyCategoryIngestionSourcebuiltIn,
		"custom":  GroupPolicyCategoryIngestionSourcecustom,
		"unknown": GroupPolicyCategoryIngestionSourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyCategoryIngestionSource(input)
	return &out, nil
}
