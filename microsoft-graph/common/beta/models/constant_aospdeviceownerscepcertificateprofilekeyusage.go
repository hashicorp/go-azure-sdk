package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileKeyUsage string

const (
	AospDeviceOwnerScepCertificateProfileKeyUsagedigitalSignature AospDeviceOwnerScepCertificateProfileKeyUsage = "DigitalSignature"
	AospDeviceOwnerScepCertificateProfileKeyUsagekeyEncipherment  AospDeviceOwnerScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileKeyUsagedigitalSignature),
		string(AospDeviceOwnerScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseAospDeviceOwnerScepCertificateProfileKeyUsage(input string) (*AospDeviceOwnerScepCertificateProfileKeyUsage, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileKeyUsage{
		"digitalsignature": AospDeviceOwnerScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  AospDeviceOwnerScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileKeyUsage(input)
	return &out, nil
}
