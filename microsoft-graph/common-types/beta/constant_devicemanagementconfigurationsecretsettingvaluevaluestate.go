package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSecretSettingValueValueState string

const (
	DeviceManagementConfigurationSecretSettingValueValueState_EncryptedValueToken DeviceManagementConfigurationSecretSettingValueValueState = "encryptedValueToken"
	DeviceManagementConfigurationSecretSettingValueValueState_Invalid             DeviceManagementConfigurationSecretSettingValueValueState = "invalid"
	DeviceManagementConfigurationSecretSettingValueValueState_NotEncrypted        DeviceManagementConfigurationSecretSettingValueValueState = "notEncrypted"
)

func PossibleValuesForDeviceManagementConfigurationSecretSettingValueValueState() []string {
	return []string{
		string(DeviceManagementConfigurationSecretSettingValueValueState_EncryptedValueToken),
		string(DeviceManagementConfigurationSecretSettingValueValueState_Invalid),
		string(DeviceManagementConfigurationSecretSettingValueValueState_NotEncrypted),
	}
}

func (s *DeviceManagementConfigurationSecretSettingValueValueState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSecretSettingValueValueState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSecretSettingValueValueState(input string) (*DeviceManagementConfigurationSecretSettingValueValueState, error) {
	vals := map[string]DeviceManagementConfigurationSecretSettingValueValueState{
		"encryptedvaluetoken": DeviceManagementConfigurationSecretSettingValueValueState_EncryptedValueToken,
		"invalid":             DeviceManagementConfigurationSecretSettingValueValueState_Invalid,
		"notencrypted":        DeviceManagementConfigurationSecretSettingValueValueState_NotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSecretSettingValueValueState(input)
	return &out, nil
}
