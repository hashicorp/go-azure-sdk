package managementgroups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupChildType string

const (
	ManagementGroupChildTypeMicrosoftPointManagementManagementGroups ManagementGroupChildType = "Microsoft.Management/managementGroups"
	ManagementGroupChildTypeSubscriptions                            ManagementGroupChildType = "/subscriptions"
)

func PossibleValuesForManagementGroupChildType() []string {
	return []string{
		string(ManagementGroupChildTypeMicrosoftPointManagementManagementGroups),
		string(ManagementGroupChildTypeSubscriptions),
	}
}

func parseManagementGroupChildType(input string) (*ManagementGroupChildType, error) {
	vals := map[string]ManagementGroupChildType{
		"microsoft.management/managementgroups": ManagementGroupChildTypeMicrosoftPointManagementManagementGroups,
		"/subscriptions":                        ManagementGroupChildTypeSubscriptions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementGroupChildType(input)
	return &out, nil
}

type ManagementGroupExpandType string

const (
	ManagementGroupExpandTypeAncestors ManagementGroupExpandType = "ancestors"
	ManagementGroupExpandTypeChildren  ManagementGroupExpandType = "children"
	ManagementGroupExpandTypePath      ManagementGroupExpandType = "path"
)

func PossibleValuesForManagementGroupExpandType() []string {
	return []string{
		string(ManagementGroupExpandTypeAncestors),
		string(ManagementGroupExpandTypeChildren),
		string(ManagementGroupExpandTypePath),
	}
}

func parseManagementGroupExpandType(input string) (*ManagementGroupExpandType, error) {
	vals := map[string]ManagementGroupExpandType{
		"ancestors": ManagementGroupExpandTypeAncestors,
		"children":  ManagementGroupExpandTypeChildren,
		"path":      ManagementGroupExpandTypePath,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementGroupExpandType(input)
	return &out, nil
}
