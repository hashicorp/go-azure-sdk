package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteEncryptionMethod string

const (
	CryptographySuiteEncryptionMethodaes128           CryptographySuiteEncryptionMethod = "Aes128"
	CryptographySuiteEncryptionMethodaes128Gcm        CryptographySuiteEncryptionMethod = "Aes128Gcm"
	CryptographySuiteEncryptionMethodaes192           CryptographySuiteEncryptionMethod = "Aes192"
	CryptographySuiteEncryptionMethodaes192Gcm        CryptographySuiteEncryptionMethod = "Aes192Gcm"
	CryptographySuiteEncryptionMethodaes256           CryptographySuiteEncryptionMethod = "Aes256"
	CryptographySuiteEncryptionMethodaes256Gcm        CryptographySuiteEncryptionMethod = "Aes256Gcm"
	CryptographySuiteEncryptionMethodchaCha20Poly1305 CryptographySuiteEncryptionMethod = "ChaCha20Poly1305"
	CryptographySuiteEncryptionMethoddes              CryptographySuiteEncryptionMethod = "Des"
	CryptographySuiteEncryptionMethodtripleDes        CryptographySuiteEncryptionMethod = "TripleDes"
)

func PossibleValuesForCryptographySuiteEncryptionMethod() []string {
	return []string{
		string(CryptographySuiteEncryptionMethodaes128),
		string(CryptographySuiteEncryptionMethodaes128Gcm),
		string(CryptographySuiteEncryptionMethodaes192),
		string(CryptographySuiteEncryptionMethodaes192Gcm),
		string(CryptographySuiteEncryptionMethodaes256),
		string(CryptographySuiteEncryptionMethodaes256Gcm),
		string(CryptographySuiteEncryptionMethodchaCha20Poly1305),
		string(CryptographySuiteEncryptionMethoddes),
		string(CryptographySuiteEncryptionMethodtripleDes),
	}
}

func parseCryptographySuiteEncryptionMethod(input string) (*CryptographySuiteEncryptionMethod, error) {
	vals := map[string]CryptographySuiteEncryptionMethod{
		"aes128":           CryptographySuiteEncryptionMethodaes128,
		"aes128gcm":        CryptographySuiteEncryptionMethodaes128Gcm,
		"aes192":           CryptographySuiteEncryptionMethodaes192,
		"aes192gcm":        CryptographySuiteEncryptionMethodaes192Gcm,
		"aes256":           CryptographySuiteEncryptionMethodaes256,
		"aes256gcm":        CryptographySuiteEncryptionMethodaes256Gcm,
		"chacha20poly1305": CryptographySuiteEncryptionMethodchaCha20Poly1305,
		"des":              CryptographySuiteEncryptionMethoddes,
		"tripledes":        CryptographySuiteEncryptionMethodtripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteEncryptionMethod(input)
	return &out, nil
}
