package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessRuleAccessLevel string

const (
	DeviceManagementExchangeAccessRuleAccessLevel_Allow      DeviceManagementExchangeAccessRuleAccessLevel = "allow"
	DeviceManagementExchangeAccessRuleAccessLevel_Block      DeviceManagementExchangeAccessRuleAccessLevel = "block"
	DeviceManagementExchangeAccessRuleAccessLevel_None       DeviceManagementExchangeAccessRuleAccessLevel = "none"
	DeviceManagementExchangeAccessRuleAccessLevel_Quarantine DeviceManagementExchangeAccessRuleAccessLevel = "quarantine"
)

func PossibleValuesForDeviceManagementExchangeAccessRuleAccessLevel() []string {
	return []string{
		string(DeviceManagementExchangeAccessRuleAccessLevel_Allow),
		string(DeviceManagementExchangeAccessRuleAccessLevel_Block),
		string(DeviceManagementExchangeAccessRuleAccessLevel_None),
		string(DeviceManagementExchangeAccessRuleAccessLevel_Quarantine),
	}
}

func (s *DeviceManagementExchangeAccessRuleAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeAccessRuleAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeAccessRuleAccessLevel(input string) (*DeviceManagementExchangeAccessRuleAccessLevel, error) {
	vals := map[string]DeviceManagementExchangeAccessRuleAccessLevel{
		"allow":      DeviceManagementExchangeAccessRuleAccessLevel_Allow,
		"block":      DeviceManagementExchangeAccessRuleAccessLevel_Block,
		"none":       DeviceManagementExchangeAccessRuleAccessLevel_None,
		"quarantine": DeviceManagementExchangeAccessRuleAccessLevel_Quarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeAccessRuleAccessLevel(input)
	return &out, nil
}
