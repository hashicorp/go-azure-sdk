package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileHashAlgorithm string

const (
	WindowsPhone81SCEPCertificateProfileHashAlgorithmsha1 WindowsPhone81SCEPCertificateProfileHashAlgorithm = "Sha1"
	WindowsPhone81SCEPCertificateProfileHashAlgorithmsha2 WindowsPhone81SCEPCertificateProfileHashAlgorithm = "Sha2"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileHashAlgorithm() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileHashAlgorithmsha1),
		string(WindowsPhone81SCEPCertificateProfileHashAlgorithmsha2),
	}
}

func parseWindowsPhone81SCEPCertificateProfileHashAlgorithm(input string) (*WindowsPhone81SCEPCertificateProfileHashAlgorithm, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileHashAlgorithm{
		"sha1": WindowsPhone81SCEPCertificateProfileHashAlgorithmsha1,
		"sha2": WindowsPhone81SCEPCertificateProfileHashAlgorithmsha2,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileHashAlgorithm(input)
	return &out, nil
}
