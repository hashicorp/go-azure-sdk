package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileHashAlgorithm string

const (
	AndroidForWorkScepCertificateProfileHashAlgorithmsha1 AndroidForWorkScepCertificateProfileHashAlgorithm = "Sha1"
	AndroidForWorkScepCertificateProfileHashAlgorithmsha2 AndroidForWorkScepCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileHashAlgorithm() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileHashAlgorithmsha1),
		string(AndroidForWorkScepCertificateProfileHashAlgorithmsha2),
	}
}

func parseAndroidForWorkScepCertificateProfileHashAlgorithm(input string) (*AndroidForWorkScepCertificateProfileHashAlgorithm, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileHashAlgorithm{
		"sha1": AndroidForWorkScepCertificateProfileHashAlgorithmsha1,
		"sha2": AndroidForWorkScepCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileHashAlgorithm(input)
	return &out, nil
}
