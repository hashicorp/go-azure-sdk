package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSettingsExternalUserLifecycleAction string

const (
	EntitlementManagementSettingsExternalUserLifecycleActionblockSignIn          EntitlementManagementSettingsExternalUserLifecycleAction = "BlockSignIn"
	EntitlementManagementSettingsExternalUserLifecycleActionblockSignInAndDelete EntitlementManagementSettingsExternalUserLifecycleAction = "BlockSignInAndDelete"
	EntitlementManagementSettingsExternalUserLifecycleActionnone                 EntitlementManagementSettingsExternalUserLifecycleAction = "None"
)

func PossibleValuesForEntitlementManagementSettingsExternalUserLifecycleAction() []string {
	return []string{
		string(EntitlementManagementSettingsExternalUserLifecycleActionblockSignIn),
		string(EntitlementManagementSettingsExternalUserLifecycleActionblockSignInAndDelete),
		string(EntitlementManagementSettingsExternalUserLifecycleActionnone),
	}
}

func parseEntitlementManagementSettingsExternalUserLifecycleAction(input string) (*EntitlementManagementSettingsExternalUserLifecycleAction, error) {
	vals := map[string]EntitlementManagementSettingsExternalUserLifecycleAction{
		"blocksignin":          EntitlementManagementSettingsExternalUserLifecycleActionblockSignIn,
		"blocksigninanddelete": EntitlementManagementSettingsExternalUserLifecycleActionblockSignInAndDelete,
		"none":                 EntitlementManagementSettingsExternalUserLifecycleActionnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntitlementManagementSettingsExternalUserLifecycleAction(input)
	return &out, nil
}
