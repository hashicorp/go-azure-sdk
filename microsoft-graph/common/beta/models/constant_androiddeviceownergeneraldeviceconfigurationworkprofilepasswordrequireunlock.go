package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockdaily             AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock = "Daily"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockdeviceDefault     AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock = "DeviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockunkownFutureValue AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock = "UnkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockdaily),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockdeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockunkownFutureValue),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock{
		"daily":             AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockdaily,
		"devicedefault":     AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockdeviceDefault,
		"unkownfuturevalue": AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlockunkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequireUnlock(input)
	return &out, nil
}
