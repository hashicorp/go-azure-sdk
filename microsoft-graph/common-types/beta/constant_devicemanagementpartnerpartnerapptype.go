package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerPartnerAppType string

const (
	DeviceManagementPartnerPartnerAppType_MultiTenantApp  DeviceManagementPartnerPartnerAppType = "multiTenantApp"
	DeviceManagementPartnerPartnerAppType_SingleTenantApp DeviceManagementPartnerPartnerAppType = "singleTenantApp"
	DeviceManagementPartnerPartnerAppType_Unknown         DeviceManagementPartnerPartnerAppType = "unknown"
)

func PossibleValuesForDeviceManagementPartnerPartnerAppType() []string {
	return []string{
		string(DeviceManagementPartnerPartnerAppType_MultiTenantApp),
		string(DeviceManagementPartnerPartnerAppType_SingleTenantApp),
		string(DeviceManagementPartnerPartnerAppType_Unknown),
	}
}

func (s *DeviceManagementPartnerPartnerAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementPartnerPartnerAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementPartnerPartnerAppType(input string) (*DeviceManagementPartnerPartnerAppType, error) {
	vals := map[string]DeviceManagementPartnerPartnerAppType{
		"multitenantapp":  DeviceManagementPartnerPartnerAppType_MultiTenantApp,
		"singletenantapp": DeviceManagementPartnerPartnerAppType_SingleTenantApp,
		"unknown":         DeviceManagementPartnerPartnerAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPartnerPartnerAppType(input)
	return &out, nil
}
