package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileHashAlgorithm string

const (
	AndroidWorkProfileScepCertificateProfileHashAlgorithmsha1 AndroidWorkProfileScepCertificateProfileHashAlgorithm = "Sha1"
	AndroidWorkProfileScepCertificateProfileHashAlgorithmsha2 AndroidWorkProfileScepCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileHashAlgorithmsha1),
		string(AndroidWorkProfileScepCertificateProfileHashAlgorithmsha2),
	}
}

func parseAndroidWorkProfileScepCertificateProfileHashAlgorithm(input string) (*AndroidWorkProfileScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileHashAlgorithm{
		"sha1": AndroidWorkProfileScepCertificateProfileHashAlgorithmsha1,
		"sha2": AndroidWorkProfileScepCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
