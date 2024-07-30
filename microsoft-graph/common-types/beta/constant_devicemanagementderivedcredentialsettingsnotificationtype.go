package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialSettingsNotificationType string

const (
	DeviceManagementDerivedCredentialSettingsNotificationType_CompanyPortal DeviceManagementDerivedCredentialSettingsNotificationType = "companyPortal"
	DeviceManagementDerivedCredentialSettingsNotificationType_Email         DeviceManagementDerivedCredentialSettingsNotificationType = "email"
	DeviceManagementDerivedCredentialSettingsNotificationType_None          DeviceManagementDerivedCredentialSettingsNotificationType = "none"
)

func PossibleValuesForDeviceManagementDerivedCredentialSettingsNotificationType() []string {
	return []string{
		string(DeviceManagementDerivedCredentialSettingsNotificationType_CompanyPortal),
		string(DeviceManagementDerivedCredentialSettingsNotificationType_Email),
		string(DeviceManagementDerivedCredentialSettingsNotificationType_None),
	}
}

func (s *DeviceManagementDerivedCredentialSettingsNotificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementDerivedCredentialSettingsNotificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementDerivedCredentialSettingsNotificationType(input string) (*DeviceManagementDerivedCredentialSettingsNotificationType, error) {
	vals := map[string]DeviceManagementDerivedCredentialSettingsNotificationType{
		"companyportal": DeviceManagementDerivedCredentialSettingsNotificationType_CompanyPortal,
		"email":         DeviceManagementDerivedCredentialSettingsNotificationType_Email,
		"none":          DeviceManagementDerivedCredentialSettingsNotificationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDerivedCredentialSettingsNotificationType(input)
	return &out, nil
}
