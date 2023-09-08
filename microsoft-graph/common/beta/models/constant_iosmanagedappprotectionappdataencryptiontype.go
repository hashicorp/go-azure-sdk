package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppDataEncryptionType string

const (
	IosManagedAppProtectionAppDataEncryptionTypeafterDeviceRestart              IosManagedAppProtectionAppDataEncryptionType = "AfterDeviceRestart"
	IosManagedAppProtectionAppDataEncryptionTypeuseDeviceSettings               IosManagedAppProtectionAppDataEncryptionType = "UseDeviceSettings"
	IosManagedAppProtectionAppDataEncryptionTypewhenDeviceLocked                IosManagedAppProtectionAppDataEncryptionType = "WhenDeviceLocked"
	IosManagedAppProtectionAppDataEncryptionTypewhenDeviceLockedExceptOpenFiles IosManagedAppProtectionAppDataEncryptionType = "WhenDeviceLockedExceptOpenFiles"
)

func PossibleValuesForIosManagedAppProtectionAppDataEncryptionType() []string {
	return []string{
		string(IosManagedAppProtectionAppDataEncryptionTypeafterDeviceRestart),
		string(IosManagedAppProtectionAppDataEncryptionTypeuseDeviceSettings),
		string(IosManagedAppProtectionAppDataEncryptionTypewhenDeviceLocked),
		string(IosManagedAppProtectionAppDataEncryptionTypewhenDeviceLockedExceptOpenFiles),
	}
}

func parseIosManagedAppProtectionAppDataEncryptionType(input string) (*IosManagedAppProtectionAppDataEncryptionType, error) {
	vals := map[string]IosManagedAppProtectionAppDataEncryptionType{
		"afterdevicerestart":              IosManagedAppProtectionAppDataEncryptionTypeafterDeviceRestart,
		"usedevicesettings":               IosManagedAppProtectionAppDataEncryptionTypeuseDeviceSettings,
		"whendevicelocked":                IosManagedAppProtectionAppDataEncryptionTypewhenDeviceLocked,
		"whendevicelockedexceptopenfiles": IosManagedAppProtectionAppDataEncryptionTypewhenDeviceLockedExceptOpenFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppDataEncryptionType(input)
	return &out, nil
}
