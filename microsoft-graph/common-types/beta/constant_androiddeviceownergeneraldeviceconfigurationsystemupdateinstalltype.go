package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Automatic     AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "automatic"
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_DeviceDefault AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "deviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Postpone      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "postpone"
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Windowed      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "windowed"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Automatic),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_DeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Postpone),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Windowed),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType{
		"automatic":     AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Automatic,
		"devicedefault": AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_DeviceDefault,
		"postpone":      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Postpone,
		"windowed":      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType_Windowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType(input)
	return &out, nil
}
