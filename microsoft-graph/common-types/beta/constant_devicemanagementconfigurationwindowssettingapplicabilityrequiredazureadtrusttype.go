package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_AddWorkAccount DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "addWorkAccount"
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_AzureAdJoined  DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "azureAdJoined"
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_MdmOnly        DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "mdmOnly"
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_None           DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "none"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_AddWorkAccount),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_AzureAdJoined),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_MdmOnly),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_None),
	}
}

func (s *DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType{
		"addworkaccount": DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_AddWorkAccount,
		"azureadjoined":  DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_AzureAdJoined,
		"mdmonly":        DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_MdmOnly,
		"none":           DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType(input)
	return &out, nil
}
