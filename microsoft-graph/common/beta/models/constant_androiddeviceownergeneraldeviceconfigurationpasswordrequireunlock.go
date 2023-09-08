package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockdaily             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock = "Daily"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockdeviceDefault     AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock = "DeviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockunkownFutureValue AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock = "UnkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockdaily),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockdeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockunkownFutureValue),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock{
		"daily":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockdaily,
		"devicedefault":     AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockdeviceDefault,
		"unkownfuturevalue": AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlockunkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequireUnlock(input)
	return &out, nil
}
