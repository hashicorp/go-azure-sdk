package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteIntegrityCheckMethod string

const (
	CryptographySuiteIntegrityCheckMethodmd5      CryptographySuiteIntegrityCheckMethod = "Md5"
	CryptographySuiteIntegrityCheckMethodsha1_160 CryptographySuiteIntegrityCheckMethod = "Sha1160"
	CryptographySuiteIntegrityCheckMethodsha1_96  CryptographySuiteIntegrityCheckMethod = "Sha196"
	CryptographySuiteIntegrityCheckMethodsha2_256 CryptographySuiteIntegrityCheckMethod = "Sha2256"
	CryptographySuiteIntegrityCheckMethodsha2_384 CryptographySuiteIntegrityCheckMethod = "Sha2384"
	CryptographySuiteIntegrityCheckMethodsha2_512 CryptographySuiteIntegrityCheckMethod = "Sha2512"
)

func PossibleValuesForCryptographySuiteIntegrityCheckMethod() []string {
	return []string{
		string(CryptographySuiteIntegrityCheckMethodmd5),
		string(CryptographySuiteIntegrityCheckMethodsha1_160),
		string(CryptographySuiteIntegrityCheckMethodsha1_96),
		string(CryptographySuiteIntegrityCheckMethodsha2_256),
		string(CryptographySuiteIntegrityCheckMethodsha2_384),
		string(CryptographySuiteIntegrityCheckMethodsha2_512),
	}
}

func parseCryptographySuiteIntegrityCheckMethod(input string) (*CryptographySuiteIntegrityCheckMethod, error) {
	vals := map[string]CryptographySuiteIntegrityCheckMethod{
		"md5":     CryptographySuiteIntegrityCheckMethodmd5,
		"sha1160": CryptographySuiteIntegrityCheckMethodsha1_160,
		"sha196":  CryptographySuiteIntegrityCheckMethodsha1_96,
		"sha2256": CryptographySuiteIntegrityCheckMethodsha2_256,
		"sha2384": CryptographySuiteIntegrityCheckMethodsha2_384,
		"sha2512": CryptographySuiteIntegrityCheckMethodsha2_512,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteIntegrityCheckMethod(input)
	return &out, nil
}
