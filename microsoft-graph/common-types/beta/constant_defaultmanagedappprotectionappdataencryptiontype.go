package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppDataEncryptionType string

const (
	DefaultManagedAppProtectionAppDataEncryptionType_AfterDeviceRestart              DefaultManagedAppProtectionAppDataEncryptionType = "afterDeviceRestart"
	DefaultManagedAppProtectionAppDataEncryptionType_UseDeviceSettings               DefaultManagedAppProtectionAppDataEncryptionType = "useDeviceSettings"
	DefaultManagedAppProtectionAppDataEncryptionType_WhenDeviceLocked                DefaultManagedAppProtectionAppDataEncryptionType = "whenDeviceLocked"
	DefaultManagedAppProtectionAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles DefaultManagedAppProtectionAppDataEncryptionType = "whenDeviceLockedExceptOpenFiles"
)

func PossibleValuesForDefaultManagedAppProtectionAppDataEncryptionType() []string {
	return []string{
		string(DefaultManagedAppProtectionAppDataEncryptionType_AfterDeviceRestart),
		string(DefaultManagedAppProtectionAppDataEncryptionType_UseDeviceSettings),
		string(DefaultManagedAppProtectionAppDataEncryptionType_WhenDeviceLocked),
		string(DefaultManagedAppProtectionAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles),
	}
}

func (s *DefaultManagedAppProtectionAppDataEncryptionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppDataEncryptionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppDataEncryptionType(input string) (*DefaultManagedAppProtectionAppDataEncryptionType, error) {
	vals := map[string]DefaultManagedAppProtectionAppDataEncryptionType{
		"afterdevicerestart":              DefaultManagedAppProtectionAppDataEncryptionType_AfterDeviceRestart,
		"usedevicesettings":               DefaultManagedAppProtectionAppDataEncryptionType_UseDeviceSettings,
		"whendevicelocked":                DefaultManagedAppProtectionAppDataEncryptionType_WhenDeviceLocked,
		"whendevicelockedexceptopenfiles": DefaultManagedAppProtectionAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppDataEncryptionType(input)
	return &out, nil
}
