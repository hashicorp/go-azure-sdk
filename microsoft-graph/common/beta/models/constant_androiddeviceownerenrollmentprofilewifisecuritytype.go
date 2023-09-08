package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileWifiSecurityType string

const (
	AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypenone AndroidDeviceOwnerEnrollmentProfileWifiSecurityType = "None"
	AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypewep  AndroidDeviceOwnerEnrollmentProfileWifiSecurityType = "Wep"
	AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypewpa  AndroidDeviceOwnerEnrollmentProfileWifiSecurityType = "Wpa"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileWifiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypenone),
		string(AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypewep),
		string(AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypewpa),
	}
}

func parseAndroidDeviceOwnerEnrollmentProfileWifiSecurityType(input string) (*AndroidDeviceOwnerEnrollmentProfileWifiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileWifiSecurityType{
		"none": AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypenone,
		"wep":  AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypewep,
		"wpa":  AndroidDeviceOwnerEnrollmentProfileWifiSecurityTypewpa,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileWifiSecurityType(input)
	return &out, nil
}
