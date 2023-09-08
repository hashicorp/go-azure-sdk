package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileKeySize string

const (
	AospDeviceOwnerScepCertificateProfileKeySizesize1024 AospDeviceOwnerScepCertificateProfileKeySize = "Size1024"
	AospDeviceOwnerScepCertificateProfileKeySizesize2048 AospDeviceOwnerScepCertificateProfileKeySize = "Size2048"
	AospDeviceOwnerScepCertificateProfileKeySizesize4096 AospDeviceOwnerScepCertificateProfileKeySize = "Size4096"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileKeySize() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileKeySizesize1024),
		string(AospDeviceOwnerScepCertificateProfileKeySizesize2048),
		string(AospDeviceOwnerScepCertificateProfileKeySizesize4096),
	}
}

func parseAospDeviceOwnerScepCertificateProfileKeySize(input string) (*AospDeviceOwnerScepCertificateProfileKeySize, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileKeySize{
		"size1024": AospDeviceOwnerScepCertificateProfileKeySizesize1024,
		"size2048": AospDeviceOwnerScepCertificateProfileKeySizesize2048,
		"size4096": AospDeviceOwnerScepCertificateProfileKeySizesize4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileKeySize(input)
	return &out, nil
}
