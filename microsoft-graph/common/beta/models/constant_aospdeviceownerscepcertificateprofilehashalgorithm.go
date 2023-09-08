package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileHashAlgorithm string

const (
	AospDeviceOwnerScepCertificateProfileHashAlgorithmsha1 AospDeviceOwnerScepCertificateProfileHashAlgorithm = "Sha1"
	AospDeviceOwnerScepCertificateProfileHashAlgorithmsha2 AospDeviceOwnerScepCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileHashAlgorithmsha1),
		string(AospDeviceOwnerScepCertificateProfileHashAlgorithmsha2),
	}
}

func parseAospDeviceOwnerScepCertificateProfileHashAlgorithm(input string) (*AospDeviceOwnerScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileHashAlgorithm{
		"sha1": AospDeviceOwnerScepCertificateProfileHashAlgorithmsha1,
		"sha2": AospDeviceOwnerScepCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
