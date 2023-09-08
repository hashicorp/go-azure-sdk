package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileKeySize string

const (
	AndroidDeviceOwnerScepCertificateProfileKeySizesize1024 AndroidDeviceOwnerScepCertificateProfileKeySize = "Size1024"
	AndroidDeviceOwnerScepCertificateProfileKeySizesize2048 AndroidDeviceOwnerScepCertificateProfileKeySize = "Size2048"
	AndroidDeviceOwnerScepCertificateProfileKeySizesize4096 AndroidDeviceOwnerScepCertificateProfileKeySize = "Size4096"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileKeySize() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileKeySizesize1024),
		string(AndroidDeviceOwnerScepCertificateProfileKeySizesize2048),
		string(AndroidDeviceOwnerScepCertificateProfileKeySizesize4096),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileKeySize(input string) (*AndroidDeviceOwnerScepCertificateProfileKeySize, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileKeySize{
		"size1024": AndroidDeviceOwnerScepCertificateProfileKeySizesize1024,
		"size2048": AndroidDeviceOwnerScepCertificateProfileKeySizesize2048,
		"size4096": AndroidDeviceOwnerScepCertificateProfileKeySizesize4096,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileKeySize(input)
	return &out, nil
}
