package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSecretSettingValueValueState string

const (
	DeviceManagementConfigurationSecretSettingValueValueStateencryptedValueToken DeviceManagementConfigurationSecretSettingValueValueState = "EncryptedValueToken"
	DeviceManagementConfigurationSecretSettingValueValueStateinvalid             DeviceManagementConfigurationSecretSettingValueValueState = "Invalid"
	DeviceManagementConfigurationSecretSettingValueValueStatenotEncrypted        DeviceManagementConfigurationSecretSettingValueValueState = "NotEncrypted"
)

func PossibleValuesForDeviceManagementConfigurationSecretSettingValueValueState() []string {
	return []string{
		string(DeviceManagementConfigurationSecretSettingValueValueStateencryptedValueToken),
		string(DeviceManagementConfigurationSecretSettingValueValueStateinvalid),
		string(DeviceManagementConfigurationSecretSettingValueValueStatenotEncrypted),
	}
}

func parseDeviceManagementConfigurationSecretSettingValueValueState(input string) (*DeviceManagementConfigurationSecretSettingValueValueState, error) {
	vals := map[string]DeviceManagementConfigurationSecretSettingValueValueState{
		"encryptedvaluetoken": DeviceManagementConfigurationSecretSettingValueValueStateencryptedValueToken,
		"invalid":             DeviceManagementConfigurationSecretSettingValueValueStateinvalid,
		"notencrypted":        DeviceManagementConfigurationSecretSettingValueValueStatenotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSecretSettingValueValueState(input)
	return &out, nil
}
