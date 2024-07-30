package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingsDerivedCredentialProvider string

const (
	DeviceManagementSettingsDerivedCredentialProvider_EntrustDataCard DeviceManagementSettingsDerivedCredentialProvider = "entrustDataCard"
	DeviceManagementSettingsDerivedCredentialProvider_Intercede       DeviceManagementSettingsDerivedCredentialProvider = "intercede"
	DeviceManagementSettingsDerivedCredentialProvider_NotConfigured   DeviceManagementSettingsDerivedCredentialProvider = "notConfigured"
	DeviceManagementSettingsDerivedCredentialProvider_Purebred        DeviceManagementSettingsDerivedCredentialProvider = "purebred"
	DeviceManagementSettingsDerivedCredentialProvider_XTec            DeviceManagementSettingsDerivedCredentialProvider = "xTec"
)

func PossibleValuesForDeviceManagementSettingsDerivedCredentialProvider() []string {
	return []string{
		string(DeviceManagementSettingsDerivedCredentialProvider_EntrustDataCard),
		string(DeviceManagementSettingsDerivedCredentialProvider_Intercede),
		string(DeviceManagementSettingsDerivedCredentialProvider_NotConfigured),
		string(DeviceManagementSettingsDerivedCredentialProvider_Purebred),
		string(DeviceManagementSettingsDerivedCredentialProvider_XTec),
	}
}

func (s *DeviceManagementSettingsDerivedCredentialProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementSettingsDerivedCredentialProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementSettingsDerivedCredentialProvider(input string) (*DeviceManagementSettingsDerivedCredentialProvider, error) {
	vals := map[string]DeviceManagementSettingsDerivedCredentialProvider{
		"entrustdatacard": DeviceManagementSettingsDerivedCredentialProvider_EntrustDataCard,
		"intercede":       DeviceManagementSettingsDerivedCredentialProvider_Intercede,
		"notconfigured":   DeviceManagementSettingsDerivedCredentialProvider_NotConfigured,
		"purebred":        DeviceManagementSettingsDerivedCredentialProvider_Purebred,
		"xtec":            DeviceManagementSettingsDerivedCredentialProvider_XTec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSettingsDerivedCredentialProvider(input)
	return &out, nil
}
