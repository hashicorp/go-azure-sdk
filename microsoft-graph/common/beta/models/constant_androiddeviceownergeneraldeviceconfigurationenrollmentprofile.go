package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilededicatedDevice AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile = "DedicatedDevice"
	AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilefullyManaged    AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile = "FullyManaged"
	AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilenotConfigured   AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile = "NotConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilededicatedDevice),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilefullyManaged),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilenotConfigured),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile{
		"dedicateddevice": AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilededicatedDevice,
		"fullymanaged":    AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilefullyManaged,
		"notconfigured":   AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfilenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationEnrollmentProfile(input)
	return &out, nil
}
