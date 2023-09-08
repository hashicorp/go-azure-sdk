package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileHashAlgorithm string

const (
	AndroidScepCertificateProfileHashAlgorithmsha1 AndroidScepCertificateProfileHashAlgorithm = "Sha1"
	AndroidScepCertificateProfileHashAlgorithmsha2 AndroidScepCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForAndroidScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidScepCertificateProfileHashAlgorithmsha1),
		string(AndroidScepCertificateProfileHashAlgorithmsha2),
	}
}

func parseAndroidScepCertificateProfileHashAlgorithm(input string) (*AndroidScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidScepCertificateProfileHashAlgorithm{
		"sha1": AndroidScepCertificateProfileHashAlgorithmsha1,
		"sha2": AndroidScepCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
