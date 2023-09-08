package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileHashAlgorithm string

const (
	SecurityFileHashAlgorithmmd5      SecurityFileHashAlgorithm = "Md5"
	SecurityFileHashAlgorithmsha1     SecurityFileHashAlgorithm = "Sha1"
	SecurityFileHashAlgorithmsha256   SecurityFileHashAlgorithm = "Sha256"
	SecurityFileHashAlgorithmsha256ac SecurityFileHashAlgorithm = "Sha256ac"
	SecurityFileHashAlgorithmunknown  SecurityFileHashAlgorithm = "Unknown"
)

func PossibleValuesForSecurityFileHashAlgorithm() []string {
	return []string{
		string(SecurityFileHashAlgorithmmd5),
		string(SecurityFileHashAlgorithmsha1),
		string(SecurityFileHashAlgorithmsha256),
		string(SecurityFileHashAlgorithmsha256ac),
		string(SecurityFileHashAlgorithmunknown),
	}
}

func parseSecurityFileHashAlgorithm(input string) (*SecurityFileHashAlgorithm, error) {
	vals := map[string]SecurityFileHashAlgorithm{
		"md5":      SecurityFileHashAlgorithmmd5,
		"sha1":     SecurityFileHashAlgorithmsha1,
		"sha256":   SecurityFileHashAlgorithmsha256,
		"sha256ac": SecurityFileHashAlgorithmsha256ac,
		"unknown":  SecurityFileHashAlgorithmunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileHashAlgorithm(input)
	return &out, nil
}
