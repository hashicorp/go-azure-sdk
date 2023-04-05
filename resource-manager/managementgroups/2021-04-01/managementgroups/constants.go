package managementgroups

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
