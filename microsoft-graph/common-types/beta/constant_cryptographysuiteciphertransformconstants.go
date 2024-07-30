package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteCipherTransformConstants string

const (
	CryptographySuiteCipherTransformConstants_Aes128           CryptographySuiteCipherTransformConstants = "aes128"
	CryptographySuiteCipherTransformConstants_Aes128Gcm        CryptographySuiteCipherTransformConstants = "aes128Gcm"
	CryptographySuiteCipherTransformConstants_Aes192           CryptographySuiteCipherTransformConstants = "aes192"
	CryptographySuiteCipherTransformConstants_Aes192Gcm        CryptographySuiteCipherTransformConstants = "aes192Gcm"
	CryptographySuiteCipherTransformConstants_Aes256           CryptographySuiteCipherTransformConstants = "aes256"
	CryptographySuiteCipherTransformConstants_Aes256Gcm        CryptographySuiteCipherTransformConstants = "aes256Gcm"
	CryptographySuiteCipherTransformConstants_ChaCha20Poly1305 CryptographySuiteCipherTransformConstants = "chaCha20Poly1305"
	CryptographySuiteCipherTransformConstants_Des              CryptographySuiteCipherTransformConstants = "des"
	CryptographySuiteCipherTransformConstants_TripleDes        CryptographySuiteCipherTransformConstants = "tripleDes"
)

func PossibleValuesForCryptographySuiteCipherTransformConstants() []string {
	return []string{
		string(CryptographySuiteCipherTransformConstants_Aes128),
		string(CryptographySuiteCipherTransformConstants_Aes128Gcm),
		string(CryptographySuiteCipherTransformConstants_Aes192),
		string(CryptographySuiteCipherTransformConstants_Aes192Gcm),
		string(CryptographySuiteCipherTransformConstants_Aes256),
		string(CryptographySuiteCipherTransformConstants_Aes256Gcm),
		string(CryptographySuiteCipherTransformConstants_ChaCha20Poly1305),
		string(CryptographySuiteCipherTransformConstants_Des),
		string(CryptographySuiteCipherTransformConstants_TripleDes),
	}
}

func (s *CryptographySuiteCipherTransformConstants) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCryptographySuiteCipherTransformConstants(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCryptographySuiteCipherTransformConstants(input string) (*CryptographySuiteCipherTransformConstants, error) {
	vals := map[string]CryptographySuiteCipherTransformConstants{
		"aes128":           CryptographySuiteCipherTransformConstants_Aes128,
		"aes128gcm":        CryptographySuiteCipherTransformConstants_Aes128Gcm,
		"aes192":           CryptographySuiteCipherTransformConstants_Aes192,
		"aes192gcm":        CryptographySuiteCipherTransformConstants_Aes192Gcm,
		"aes256":           CryptographySuiteCipherTransformConstants_Aes256,
		"aes256gcm":        CryptographySuiteCipherTransformConstants_Aes256Gcm,
		"chacha20poly1305": CryptographySuiteCipherTransformConstants_ChaCha20Poly1305,
		"des":              CryptographySuiteCipherTransformConstants_Des,
		"tripledes":        CryptographySuiteCipherTransformConstants_TripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteCipherTransformConstants(input)
	return &out, nil
}
