package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteEncryptionMethod string

const (
	CryptographySuiteEncryptionMethod_Aes128           CryptographySuiteEncryptionMethod = "aes128"
	CryptographySuiteEncryptionMethod_Aes128Gcm        CryptographySuiteEncryptionMethod = "aes128Gcm"
	CryptographySuiteEncryptionMethod_Aes192           CryptographySuiteEncryptionMethod = "aes192"
	CryptographySuiteEncryptionMethod_Aes192Gcm        CryptographySuiteEncryptionMethod = "aes192Gcm"
	CryptographySuiteEncryptionMethod_Aes256           CryptographySuiteEncryptionMethod = "aes256"
	CryptographySuiteEncryptionMethod_Aes256Gcm        CryptographySuiteEncryptionMethod = "aes256Gcm"
	CryptographySuiteEncryptionMethod_ChaCha20Poly1305 CryptographySuiteEncryptionMethod = "chaCha20Poly1305"
	CryptographySuiteEncryptionMethod_Des              CryptographySuiteEncryptionMethod = "des"
	CryptographySuiteEncryptionMethod_TripleDes        CryptographySuiteEncryptionMethod = "tripleDes"
)

func PossibleValuesForCryptographySuiteEncryptionMethod() []string {
	return []string{
		string(CryptographySuiteEncryptionMethod_Aes128),
		string(CryptographySuiteEncryptionMethod_Aes128Gcm),
		string(CryptographySuiteEncryptionMethod_Aes192),
		string(CryptographySuiteEncryptionMethod_Aes192Gcm),
		string(CryptographySuiteEncryptionMethod_Aes256),
		string(CryptographySuiteEncryptionMethod_Aes256Gcm),
		string(CryptographySuiteEncryptionMethod_ChaCha20Poly1305),
		string(CryptographySuiteEncryptionMethod_Des),
		string(CryptographySuiteEncryptionMethod_TripleDes),
	}
}

func (s *CryptographySuiteEncryptionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCryptographySuiteEncryptionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCryptographySuiteEncryptionMethod(input string) (*CryptographySuiteEncryptionMethod, error) {
	vals := map[string]CryptographySuiteEncryptionMethod{
		"aes128":           CryptographySuiteEncryptionMethod_Aes128,
		"aes128gcm":        CryptographySuiteEncryptionMethod_Aes128Gcm,
		"aes192":           CryptographySuiteEncryptionMethod_Aes192,
		"aes192gcm":        CryptographySuiteEncryptionMethod_Aes192Gcm,
		"aes256":           CryptographySuiteEncryptionMethod_Aes256,
		"aes256gcm":        CryptographySuiteEncryptionMethod_Aes256Gcm,
		"chacha20poly1305": CryptographySuiteEncryptionMethod_ChaCha20Poly1305,
		"des":              CryptographySuiteEncryptionMethod_Des,
		"tripledes":        CryptographySuiteEncryptionMethod_TripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteEncryptionMethod(input)
	return &out, nil
}
