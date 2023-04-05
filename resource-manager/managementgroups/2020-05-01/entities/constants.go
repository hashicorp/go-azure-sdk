package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InheritedPermissions string

const (
	InheritedPermissionsDelete   InheritedPermissions = "delete"
	InheritedPermissionsEdit     InheritedPermissions = "edit"
	InheritedPermissionsNoaccess InheritedPermissions = "noaccess"
	InheritedPermissionsView     InheritedPermissions = "view"
)

func PossibleValuesForInheritedPermissions() []string {
	return []string{
		string(InheritedPermissionsDelete),
		string(InheritedPermissionsEdit),
		string(InheritedPermissionsNoaccess),
		string(InheritedPermissionsView),
	}
}

type Permissions string

const (
	PermissionsDelete   Permissions = "delete"
	PermissionsEdit     Permissions = "edit"
	PermissionsNoaccess Permissions = "noaccess"
	PermissionsView     Permissions = "view"
)

func PossibleValuesForPermissions() []string {
	return []string{
		string(PermissionsDelete),
		string(PermissionsEdit),
		string(PermissionsNoaccess),
		string(PermissionsView),
	}
}

type Search string

const (
	SearchAllowedChildren             Search = "AllowedChildren"
	SearchAllowedParents              Search = "AllowedParents"
	SearchChildrenOnly                Search = "ChildrenOnly"
	SearchParentAndFirstLevelChildren Search = "ParentAndFirstLevelChildren"
	SearchParentOnly                  Search = "ParentOnly"
)

func PossibleValuesForSearch() []string {
	return []string{
		string(SearchAllowedChildren),
		string(SearchAllowedParents),
		string(SearchChildrenOnly),
		string(SearchParentAndFirstLevelChildren),
		string(SearchParentOnly),
	}
}

type View string

const (
	ViewAudit             View = "Audit"
	ViewFullHierarchy     View = "FullHierarchy"
	ViewGroupsOnly        View = "GroupsOnly"
	ViewSubscriptionsOnly View = "SubscriptionsOnly"
)

func PossibleValuesForView() []string {
	return []string{
		string(ViewAudit),
		string(ViewFullHierarchy),
		string(ViewGroupsOnly),
		string(ViewSubscriptionsOnly),
	}
}
