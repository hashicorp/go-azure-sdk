package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteAuthenticationTransformConstants string

const (
	CryptographySuiteAuthenticationTransformConstantsaes128Gcm   CryptographySuiteAuthenticationTransformConstants = "Aes128Gcm"
	CryptographySuiteAuthenticationTransformConstantsaes192Gcm   CryptographySuiteAuthenticationTransformConstants = "Aes192Gcm"
	CryptographySuiteAuthenticationTransformConstantsaes256Gcm   CryptographySuiteAuthenticationTransformConstants = "Aes256Gcm"
	CryptographySuiteAuthenticationTransformConstantsmd5_96      CryptographySuiteAuthenticationTransformConstants = "Md596"
	CryptographySuiteAuthenticationTransformConstantssha1_96     CryptographySuiteAuthenticationTransformConstants = "Sha196"
	CryptographySuiteAuthenticationTransformConstantssha_256_128 CryptographySuiteAuthenticationTransformConstants = "Sha256128"
)

func PossibleValuesForCryptographySuiteAuthenticationTransformConstants() []string {
	return []string{
		string(CryptographySuiteAuthenticationTransformConstantsaes128Gcm),
		string(CryptographySuiteAuthenticationTransformConstantsaes192Gcm),
		string(CryptographySuiteAuthenticationTransformConstantsaes256Gcm),
		string(CryptographySuiteAuthenticationTransformConstantsmd5_96),
		string(CryptographySuiteAuthenticationTransformConstantssha1_96),
		string(CryptographySuiteAuthenticationTransformConstantssha_256_128),
	}
}

func parseCryptographySuiteAuthenticationTransformConstants(input string) (*CryptographySuiteAuthenticationTransformConstants, error) {
	vals := map[string]CryptographySuiteAuthenticationTransformConstants{
		"aes128gcm": CryptographySuiteAuthenticationTransformConstantsaes128Gcm,
		"aes192gcm": CryptographySuiteAuthenticationTransformConstantsaes192Gcm,
		"aes256gcm": CryptographySuiteAuthenticationTransformConstantsaes256Gcm,
		"md596":     CryptographySuiteAuthenticationTransformConstantsmd5_96,
		"sha196":    CryptographySuiteAuthenticationTransformConstantssha1_96,
		"sha256128": CryptographySuiteAuthenticationTransformConstantssha_256_128,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteAuthenticationTransformConstants(input)
	return &out, nil
}
