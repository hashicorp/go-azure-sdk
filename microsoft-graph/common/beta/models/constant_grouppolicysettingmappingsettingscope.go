package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingSettingScope string

const (
	GroupPolicySettingMappingSettingScopedevice  GroupPolicySettingMappingSettingScope = "Device"
	GroupPolicySettingMappingSettingScopeunknown GroupPolicySettingMappingSettingScope = "Unknown"
	GroupPolicySettingMappingSettingScopeuser    GroupPolicySettingMappingSettingScope = "User"
)

func PossibleValuesForGroupPolicySettingMappingSettingScope() []string {
	return []string{
		string(GroupPolicySettingMappingSettingScopedevice),
		string(GroupPolicySettingMappingSettingScopeunknown),
		string(GroupPolicySettingMappingSettingScopeuser),
	}
}

func parseGroupPolicySettingMappingSettingScope(input string) (*GroupPolicySettingMappingSettingScope, error) {
	vals := map[string]GroupPolicySettingMappingSettingScope{
		"device":  GroupPolicySettingMappingSettingScopedevice,
		"unknown": GroupPolicySettingMappingSettingScopeunknown,
		"user":    GroupPolicySettingMappingSettingScopeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingMappingSettingScope(input)
	return &out, nil
}
