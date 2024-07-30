package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnsupportedGroupPolicyExtensionSettingScope string

const (
	UnsupportedGroupPolicyExtensionSettingScope_Device  UnsupportedGroupPolicyExtensionSettingScope = "device"
	UnsupportedGroupPolicyExtensionSettingScope_Unknown UnsupportedGroupPolicyExtensionSettingScope = "unknown"
	UnsupportedGroupPolicyExtensionSettingScope_User    UnsupportedGroupPolicyExtensionSettingScope = "user"
)

func PossibleValuesForUnsupportedGroupPolicyExtensionSettingScope() []string {
	return []string{
		string(UnsupportedGroupPolicyExtensionSettingScope_Device),
		string(UnsupportedGroupPolicyExtensionSettingScope_Unknown),
		string(UnsupportedGroupPolicyExtensionSettingScope_User),
	}
}

func (s *UnsupportedGroupPolicyExtensionSettingScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnsupportedGroupPolicyExtensionSettingScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnsupportedGroupPolicyExtensionSettingScope(input string) (*UnsupportedGroupPolicyExtensionSettingScope, error) {
	vals := map[string]UnsupportedGroupPolicyExtensionSettingScope{
		"device":  UnsupportedGroupPolicyExtensionSettingScope_Device,
		"unknown": UnsupportedGroupPolicyExtensionSettingScope_Unknown,
		"user":    UnsupportedGroupPolicyExtensionSettingScope_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnsupportedGroupPolicyExtensionSettingScope(input)
	return &out, nil
}
