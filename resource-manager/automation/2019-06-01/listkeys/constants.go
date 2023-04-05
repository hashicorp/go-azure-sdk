package listkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationKeyName string

const (
	AutomationKeyNamePrimary   AutomationKeyName = "Primary"
	AutomationKeyNameSecondary AutomationKeyName = "Secondary"
)

func PossibleValuesForAutomationKeyName() []string {
	return []string{
		string(AutomationKeyNamePrimary),
		string(AutomationKeyNameSecondary),
	}
}

type AutomationKeyPermissions string

const (
	AutomationKeyPermissionsFull AutomationKeyPermissions = "Full"
	AutomationKeyPermissionsRead AutomationKeyPermissions = "Read"
)

func PossibleValuesForAutomationKeyPermissions() []string {
	return []string{
		string(AutomationKeyPermissionsFull),
		string(AutomationKeyPermissionsRead),
	}
}
