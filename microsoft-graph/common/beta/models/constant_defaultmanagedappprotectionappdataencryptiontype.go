package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppDataEncryptionType string

const (
	DefaultManagedAppProtectionAppDataEncryptionTypeafterDeviceRestart              DefaultManagedAppProtectionAppDataEncryptionType = "AfterDeviceRestart"
	DefaultManagedAppProtectionAppDataEncryptionTypeuseDeviceSettings               DefaultManagedAppProtectionAppDataEncryptionType = "UseDeviceSettings"
	DefaultManagedAppProtectionAppDataEncryptionTypewhenDeviceLocked                DefaultManagedAppProtectionAppDataEncryptionType = "WhenDeviceLocked"
	DefaultManagedAppProtectionAppDataEncryptionTypewhenDeviceLockedExceptOpenFiles DefaultManagedAppProtectionAppDataEncryptionType = "WhenDeviceLockedExceptOpenFiles"
)

func PossibleValuesForDefaultManagedAppProtectionAppDataEncryptionType() []string {
	return []string{
		string(DefaultManagedAppProtectionAppDataEncryptionTypeafterDeviceRestart),
		string(DefaultManagedAppProtectionAppDataEncryptionTypeuseDeviceSettings),
		string(DefaultManagedAppProtectionAppDataEncryptionTypewhenDeviceLocked),
		string(DefaultManagedAppProtectionAppDataEncryptionTypewhenDeviceLockedExceptOpenFiles),
	}
}

func parseDefaultManagedAppProtectionAppDataEncryptionType(input string) (*DefaultManagedAppProtectionAppDataEncryptionType, error) {
	vals := map[string]DefaultManagedAppProtectionAppDataEncryptionType{
		"afterdevicerestart":              DefaultManagedAppProtectionAppDataEncryptionTypeafterDeviceRestart,
		"usedevicesettings":               DefaultManagedAppProtectionAppDataEncryptionTypeuseDeviceSettings,
		"whendevicelocked":                DefaultManagedAppProtectionAppDataEncryptionTypewhenDeviceLocked,
		"whendevicelockedexceptopenfiles": DefaultManagedAppProtectionAppDataEncryptionTypewhenDeviceLockedExceptOpenFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppDataEncryptionType(input)
	return &out, nil
}
