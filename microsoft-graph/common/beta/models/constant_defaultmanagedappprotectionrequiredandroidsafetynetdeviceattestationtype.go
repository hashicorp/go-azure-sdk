package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType string

const (
	DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrity                       DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "BasicIntegrity"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrityAndDeviceCertification DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "BasicIntegrityAndDeviceCertification"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypenone                                 DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType = "None"
)

func PossibleValuesForDefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType() []string {
	return []string{
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrity),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrityAndDeviceCertification),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypenone),
	}
}

func parseDefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input string) (*DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType, error) {
	vals := map[string]DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType{
		"basicintegrity":                       DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrity,
		"basicintegrityanddevicecertification": DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypebasicIntegrityAndDeviceCertification,
		"none":                                 DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionRequiredAndroidSafetyNetDeviceAttestationType(input)
	return &out, nil
}
