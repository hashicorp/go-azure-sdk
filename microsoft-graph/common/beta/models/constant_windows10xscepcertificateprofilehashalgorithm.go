package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileHashAlgorithm string

const (
	Windows10XSCEPCertificateProfileHashAlgorithmsha1 Windows10XSCEPCertificateProfileHashAlgorithm = "Sha1"
	Windows10XSCEPCertificateProfileHashAlgorithmsha2 Windows10XSCEPCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForWindows10XSCEPCertificateProfileHashAlgorithm() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileHashAlgorithmsha1),
		string(Windows10XSCEPCertificateProfileHashAlgorithmsha2),
	}
}

func parseWindows10XSCEPCertificateProfileHashAlgorithm(input string) (*Windows10XSCEPCertificateProfileHashAlgorithm, error) {
	vals := map[string]Windows10XSCEPCertificateProfileHashAlgorithm{
		"sha1": Windows10XSCEPCertificateProfileHashAlgorithmsha1,
		"sha2": Windows10XSCEPCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileHashAlgorithm(input)
	return &out, nil
}
