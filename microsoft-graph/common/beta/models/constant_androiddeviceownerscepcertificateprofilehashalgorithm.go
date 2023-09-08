package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileHashAlgorithm string

const (
	AndroidDeviceOwnerScepCertificateProfileHashAlgorithmsha1 AndroidDeviceOwnerScepCertificateProfileHashAlgorithm = "Sha1"
	AndroidDeviceOwnerScepCertificateProfileHashAlgorithmsha2 AndroidDeviceOwnerScepCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileHashAlgorithmsha1),
		string(AndroidDeviceOwnerScepCertificateProfileHashAlgorithmsha2),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileHashAlgorithm(input string) (*AndroidDeviceOwnerScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileHashAlgorithm{
		"sha1": AndroidDeviceOwnerScepCertificateProfileHashAlgorithmsha1,
		"sha2": AndroidDeviceOwnerScepCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
