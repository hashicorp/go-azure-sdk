package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel string

const (
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Allow      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "allow"
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Block      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "block"
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_None       DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "none"
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Quarantine DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "quarantine"
)

func PossibleValuesForDeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel() []string {
	return []string{
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Allow),
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Block),
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_None),
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Quarantine),
	}
}

func (s *DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel(input string) (*DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel, error) {
	vals := map[string]DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel{
		"allow":      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Allow,
		"block":      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Block,
		"none":       DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_None,
		"quarantine": DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel_Quarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel(input)
	return &out, nil
}
