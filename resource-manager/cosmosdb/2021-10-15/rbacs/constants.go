package rbacs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleDefinitionType string

const (
	RoleDefinitionTypeBuiltInRole RoleDefinitionType = "BuiltInRole"
	RoleDefinitionTypeCustomRole  RoleDefinitionType = "CustomRole"
)

func PossibleValuesForRoleDefinitionType() []string {
	return []string{
		string(RoleDefinitionTypeBuiltInRole),
		string(RoleDefinitionTypeCustomRole),
	}
}
