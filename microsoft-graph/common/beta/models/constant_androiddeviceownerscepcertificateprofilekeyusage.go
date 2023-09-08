package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileKeyUsage string

const (
	AndroidDeviceOwnerScepCertificateProfileKeyUsagedigitalSignature AndroidDeviceOwnerScepCertificateProfileKeyUsage = "DigitalSignature"
	AndroidDeviceOwnerScepCertificateProfileKeyUsagekeyEncipherment  AndroidDeviceOwnerScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileKeyUsagedigitalSignature),
		string(AndroidDeviceOwnerScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileKeyUsage(input string) (*AndroidDeviceOwnerScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidDeviceOwnerScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  AndroidDeviceOwnerScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileKeyUsage(input)
	return &out, nil
}
