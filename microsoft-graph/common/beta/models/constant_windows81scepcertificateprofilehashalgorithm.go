package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileHashAlgorithm string

const (
	Windows81SCEPCertificateProfileHashAlgorithmsha1 Windows81SCEPCertificateProfileHashAlgorithm = "Sha1"
	Windows81SCEPCertificateProfileHashAlgorithmsha2 Windows81SCEPCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForWindows81SCEPCertificateProfileHashAlgorithm() []string {
	return []string{
		string(Windows81SCEPCertificateProfileHashAlgorithmsha1),
		string(Windows81SCEPCertificateProfileHashAlgorithmsha2),
	}
}

func parseWindows81SCEPCertificateProfileHashAlgorithm(input string) (*Windows81SCEPCertificateProfileHashAlgorithm, error) {
	vals := map[string]Windows81SCEPCertificateProfileHashAlgorithm{
		"sha1": Windows81SCEPCertificateProfileHashAlgorithmsha1,
		"sha2": Windows81SCEPCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileHashAlgorithm(input)
	return &out, nil
}
