package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_AutoDeny      AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "autoDeny"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_AutoGrant     AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "autoGrant"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_DeviceDefault AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "deviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_Prompt        AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy = "prompt"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_AutoDeny),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_AutoGrant),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_DeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_Prompt),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy{
		"autodeny":      AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_AutoDeny,
		"autogrant":     AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_AutoGrant,
		"devicedefault": AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_DeviceDefault,
		"prompt":        AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationAppsDefaultPermissionPolicy(input)
	return &out, nil
}
