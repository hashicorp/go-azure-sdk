package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType string

const (
	AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrity                       AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "BasicIntegrity"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrityAndDeviceCertification AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "BasicIntegrityAndDeviceCertification"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypenone                                 AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "None"
)

func PossibleValuesForAndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType() []string {
	return []string{
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrity),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrityAndDeviceCertification),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypenone),
	}
}

func parseAndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input string) (*AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType, error) {
	vals := map[string]AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType{
		"basicintegrity":                       AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrity,
		"basicintegrityanddevicecertification": AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrityAndDeviceCertification,
		"none":                                 AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input)
	return &out, nil
}
