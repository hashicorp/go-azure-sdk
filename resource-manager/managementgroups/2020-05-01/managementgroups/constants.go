package managementgroups

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
