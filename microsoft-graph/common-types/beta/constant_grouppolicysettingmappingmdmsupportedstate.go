package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingMdmSupportedState string

const (
	GroupPolicySettingMappingMdmSupportedState_Deprecated  GroupPolicySettingMappingMdmSupportedState = "deprecated"
	GroupPolicySettingMappingMdmSupportedState_Supported   GroupPolicySettingMappingMdmSupportedState = "supported"
	GroupPolicySettingMappingMdmSupportedState_Unknown     GroupPolicySettingMappingMdmSupportedState = "unknown"
	GroupPolicySettingMappingMdmSupportedState_Unsupported GroupPolicySettingMappingMdmSupportedState = "unsupported"
)

func PossibleValuesForGroupPolicySettingMappingMdmSupportedState() []string {
	return []string{
		string(GroupPolicySettingMappingMdmSupportedState_Deprecated),
		string(GroupPolicySettingMappingMdmSupportedState_Supported),
		string(GroupPolicySettingMappingMdmSupportedState_Unknown),
		string(GroupPolicySettingMappingMdmSupportedState_Unsupported),
	}
}

func (s *GroupPolicySettingMappingMdmSupportedState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicySettingMappingMdmSupportedState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicySettingMappingMdmSupportedState(input string) (*GroupPolicySettingMappingMdmSupportedState, error) {
	vals := map[string]GroupPolicySettingMappingMdmSupportedState{
		"deprecated":  GroupPolicySettingMappingMdmSupportedState_Deprecated,
		"supported":   GroupPolicySettingMappingMdmSupportedState_Supported,
		"unknown":     GroupPolicySettingMappingMdmSupportedState_Unknown,
		"unsupported": GroupPolicySettingMappingMdmSupportedState_Unsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingMappingMdmSupportedState(input)
	return &out, nil
}
