package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileHashAlgorithm string

const (
	MacOSScepCertificateProfileHashAlgorithmsha1 MacOSScepCertificateProfileHashAlgorithm = "Sha1"
	MacOSScepCertificateProfileHashAlgorithmsha2 MacOSScepCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForMacOSScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(MacOSScepCertificateProfileHashAlgorithmsha1),
		string(MacOSScepCertificateProfileHashAlgorithmsha2),
	}
}

func parseMacOSScepCertificateProfileHashAlgorithm(input string) (*MacOSScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]MacOSScepCertificateProfileHashAlgorithm{
		"sha1": MacOSScepCertificateProfileHashAlgorithmsha1,
		"sha2": MacOSScepCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
