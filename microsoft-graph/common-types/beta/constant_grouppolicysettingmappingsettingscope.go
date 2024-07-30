package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingSettingScope string

const (
	GroupPolicySettingMappingSettingScope_Device  GroupPolicySettingMappingSettingScope = "device"
	GroupPolicySettingMappingSettingScope_Unknown GroupPolicySettingMappingSettingScope = "unknown"
	GroupPolicySettingMappingSettingScope_User    GroupPolicySettingMappingSettingScope = "user"
)

func PossibleValuesForGroupPolicySettingMappingSettingScope() []string {
	return []string{
		string(GroupPolicySettingMappingSettingScope_Device),
		string(GroupPolicySettingMappingSettingScope_Unknown),
		string(GroupPolicySettingMappingSettingScope_User),
	}
}

func (s *GroupPolicySettingMappingSettingScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicySettingMappingSettingScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicySettingMappingSettingScope(input string) (*GroupPolicySettingMappingSettingScope, error) {
	vals := map[string]GroupPolicySettingMappingSettingScope{
		"device":  GroupPolicySettingMappingSettingScope_Device,
		"unknown": GroupPolicySettingMappingSettingScope_Unknown,
		"user":    GroupPolicySettingMappingSettingScope_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingMappingSettingScope(input)
	return &out, nil
}
