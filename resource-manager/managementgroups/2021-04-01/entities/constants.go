package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitySearchType string

const (
	EntitySearchTypeAllowedChildren             EntitySearchType = "AllowedChildren"
	EntitySearchTypeAllowedParents              EntitySearchType = "AllowedParents"
	EntitySearchTypeChildrenOnly                EntitySearchType = "ChildrenOnly"
	EntitySearchTypeParentAndFirstLevelChildren EntitySearchType = "ParentAndFirstLevelChildren"
	EntitySearchTypeParentOnly                  EntitySearchType = "ParentOnly"
)

func PossibleValuesForEntitySearchType() []string {
	return []string{
		string(EntitySearchTypeAllowedChildren),
		string(EntitySearchTypeAllowedParents),
		string(EntitySearchTypeChildrenOnly),
		string(EntitySearchTypeParentAndFirstLevelChildren),
		string(EntitySearchTypeParentOnly),
	}
}

type EntityViewParameterType string

const (
	EntityViewParameterTypeAudit             EntityViewParameterType = "Audit"
	EntityViewParameterTypeFullHierarchy     EntityViewParameterType = "FullHierarchy"
	EntityViewParameterTypeGroupsOnly        EntityViewParameterType = "GroupsOnly"
	EntityViewParameterTypeSubscriptionsOnly EntityViewParameterType = "SubscriptionsOnly"
)

func PossibleValuesForEntityViewParameterType() []string {
	return []string{
		string(EntityViewParameterTypeAudit),
		string(EntityViewParameterTypeFullHierarchy),
		string(EntityViewParameterTypeGroupsOnly),
		string(EntityViewParameterTypeSubscriptionsOnly),
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
