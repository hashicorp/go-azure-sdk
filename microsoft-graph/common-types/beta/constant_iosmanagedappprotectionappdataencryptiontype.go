package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppDataEncryptionType string

const (
	IosManagedAppProtectionAppDataEncryptionType_AfterDeviceRestart              IosManagedAppProtectionAppDataEncryptionType = "afterDeviceRestart"
	IosManagedAppProtectionAppDataEncryptionType_UseDeviceSettings               IosManagedAppProtectionAppDataEncryptionType = "useDeviceSettings"
	IosManagedAppProtectionAppDataEncryptionType_WhenDeviceLocked                IosManagedAppProtectionAppDataEncryptionType = "whenDeviceLocked"
	IosManagedAppProtectionAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles IosManagedAppProtectionAppDataEncryptionType = "whenDeviceLockedExceptOpenFiles"
)

func PossibleValuesForIosManagedAppProtectionAppDataEncryptionType() []string {
	return []string{
		string(IosManagedAppProtectionAppDataEncryptionType_AfterDeviceRestart),
		string(IosManagedAppProtectionAppDataEncryptionType_UseDeviceSettings),
		string(IosManagedAppProtectionAppDataEncryptionType_WhenDeviceLocked),
		string(IosManagedAppProtectionAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles),
	}
}

func (s *IosManagedAppProtectionAppDataEncryptionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAppDataEncryptionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAppDataEncryptionType(input string) (*IosManagedAppProtectionAppDataEncryptionType, error) {
	vals := map[string]IosManagedAppProtectionAppDataEncryptionType{
		"afterdevicerestart":              IosManagedAppProtectionAppDataEncryptionType_AfterDeviceRestart,
		"usedevicesettings":               IosManagedAppProtectionAppDataEncryptionType_UseDeviceSettings,
		"whendevicelocked":                IosManagedAppProtectionAppDataEncryptionType_WhenDeviceLocked,
		"whendevicelockedexceptopenfiles": IosManagedAppProtectionAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppDataEncryptionType(input)
	return &out, nil
}
