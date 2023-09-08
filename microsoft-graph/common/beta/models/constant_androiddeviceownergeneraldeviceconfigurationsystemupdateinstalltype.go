package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypeautomatic     AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "Automatic"
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypedeviceDefault AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "DeviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypepostpone      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "Postpone"
	AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypewindowed      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType = "Windowed"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypeautomatic),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypedeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypepostpone),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypewindowed),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType{
		"automatic":     AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypeautomatic,
		"devicedefault": AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypedeviceDefault,
		"postpone":      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypepostpone,
		"windowed":      AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallTypewindowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationSystemUpdateInstallType(input)
	return &out, nil
}
