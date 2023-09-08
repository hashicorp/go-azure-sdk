package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingMdmSupportedState string

const (
	GroupPolicySettingMappingMdmSupportedStatedeprecated  GroupPolicySettingMappingMdmSupportedState = "Deprecated"
	GroupPolicySettingMappingMdmSupportedStatesupported   GroupPolicySettingMappingMdmSupportedState = "Supported"
	GroupPolicySettingMappingMdmSupportedStateunknown     GroupPolicySettingMappingMdmSupportedState = "Unknown"
	GroupPolicySettingMappingMdmSupportedStateunsupported GroupPolicySettingMappingMdmSupportedState = "Unsupported"
)

func PossibleValuesForGroupPolicySettingMappingMdmSupportedState() []string {
	return []string{
		string(GroupPolicySettingMappingMdmSupportedStatedeprecated),
		string(GroupPolicySettingMappingMdmSupportedStatesupported),
		string(GroupPolicySettingMappingMdmSupportedStateunknown),
		string(GroupPolicySettingMappingMdmSupportedStateunsupported),
	}
}

func parseGroupPolicySettingMappingMdmSupportedState(input string) (*GroupPolicySettingMappingMdmSupportedState, error) {
	vals := map[string]GroupPolicySettingMappingMdmSupportedState{
		"deprecated":  GroupPolicySettingMappingMdmSupportedStatedeprecated,
		"supported":   GroupPolicySettingMappingMdmSupportedStatesupported,
		"unknown":     GroupPolicySettingMappingMdmSupportedStateunknown,
		"unsupported": GroupPolicySettingMappingMdmSupportedStateunsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingMappingMdmSupportedState(input)
	return &out, nil
}
