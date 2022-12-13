package managementgroups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Expand string

const (
	ExpandChildren Expand = "children"
	ExpandPath     Expand = "path"
)

func PossibleValuesForExpand() []string {
	return []string{
		string(ExpandChildren),
		string(ExpandPath),
	}
}

func parseExpand(input string) (*Expand, error) {
	vals := map[string]Expand{
		"children": ExpandChildren,
		"path":     ExpandPath,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Expand(input)
	return &out, nil
}

type Type string

const (
	TypeMicrosoftPointManagementManagementGroups Type = "Microsoft.Management/managementGroups"
	TypeSubscriptions                            Type = "/subscriptions"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointManagementManagementGroups),
		string(TypeSubscriptions),
	}
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"microsoft.management/managementgroups": TypeMicrosoftPointManagementManagementGroups,
		"/subscriptions":                        TypeSubscriptions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
