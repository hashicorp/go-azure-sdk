package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteCipherTransformConstants string

const (
	CryptographySuiteCipherTransformConstantsaes128           CryptographySuiteCipherTransformConstants = "Aes128"
	CryptographySuiteCipherTransformConstantsaes128Gcm        CryptographySuiteCipherTransformConstants = "Aes128Gcm"
	CryptographySuiteCipherTransformConstantsaes192           CryptographySuiteCipherTransformConstants = "Aes192"
	CryptographySuiteCipherTransformConstantsaes192Gcm        CryptographySuiteCipherTransformConstants = "Aes192Gcm"
	CryptographySuiteCipherTransformConstantsaes256           CryptographySuiteCipherTransformConstants = "Aes256"
	CryptographySuiteCipherTransformConstantsaes256Gcm        CryptographySuiteCipherTransformConstants = "Aes256Gcm"
	CryptographySuiteCipherTransformConstantschaCha20Poly1305 CryptographySuiteCipherTransformConstants = "ChaCha20Poly1305"
	CryptographySuiteCipherTransformConstantsdes              CryptographySuiteCipherTransformConstants = "Des"
	CryptographySuiteCipherTransformConstantstripleDes        CryptographySuiteCipherTransformConstants = "TripleDes"
)

func PossibleValuesForCryptographySuiteCipherTransformConstants() []string {
	return []string{
		string(CryptographySuiteCipherTransformConstantsaes128),
		string(CryptographySuiteCipherTransformConstantsaes128Gcm),
		string(CryptographySuiteCipherTransformConstantsaes192),
		string(CryptographySuiteCipherTransformConstantsaes192Gcm),
		string(CryptographySuiteCipherTransformConstantsaes256),
		string(CryptographySuiteCipherTransformConstantsaes256Gcm),
		string(CryptographySuiteCipherTransformConstantschaCha20Poly1305),
		string(CryptographySuiteCipherTransformConstantsdes),
		string(CryptographySuiteCipherTransformConstantstripleDes),
	}
}

func parseCryptographySuiteCipherTransformConstants(input string) (*CryptographySuiteCipherTransformConstants, error) {
	vals := map[string]CryptographySuiteCipherTransformConstants{
		"aes128":           CryptographySuiteCipherTransformConstantsaes128,
		"aes128gcm":        CryptographySuiteCipherTransformConstantsaes128Gcm,
		"aes192":           CryptographySuiteCipherTransformConstantsaes192,
		"aes192gcm":        CryptographySuiteCipherTransformConstantsaes192Gcm,
		"aes256":           CryptographySuiteCipherTransformConstantsaes256,
		"aes256gcm":        CryptographySuiteCipherTransformConstantsaes256Gcm,
		"chacha20poly1305": CryptographySuiteCipherTransformConstantschaCha20Poly1305,
		"des":              CryptographySuiteCipherTransformConstantsdes,
		"tripledes":        CryptographySuiteCipherTransformConstantstripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteCipherTransformConstants(input)
	return &out, nil
}
