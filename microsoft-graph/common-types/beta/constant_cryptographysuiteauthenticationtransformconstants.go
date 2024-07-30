package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteAuthenticationTransformConstants string

const (
	CryptographySuiteAuthenticationTransformConstants_Aes128Gcm CryptographySuiteAuthenticationTransformConstants = "aes128Gcm"
	CryptographySuiteAuthenticationTransformConstants_Aes192Gcm CryptographySuiteAuthenticationTransformConstants = "aes192Gcm"
	CryptographySuiteAuthenticationTransformConstants_Aes256Gcm CryptographySuiteAuthenticationTransformConstants = "aes256Gcm"
	CryptographySuiteAuthenticationTransformConstants_Md596     CryptographySuiteAuthenticationTransformConstants = "md5_96"
	CryptographySuiteAuthenticationTransformConstants_Sha196    CryptographySuiteAuthenticationTransformConstants = "sha1_96"
	CryptographySuiteAuthenticationTransformConstants_Sha256128 CryptographySuiteAuthenticationTransformConstants = "sha_256_128"
)

func PossibleValuesForCryptographySuiteAuthenticationTransformConstants() []string {
	return []string{
		string(CryptographySuiteAuthenticationTransformConstants_Aes128Gcm),
		string(CryptographySuiteAuthenticationTransformConstants_Aes192Gcm),
		string(CryptographySuiteAuthenticationTransformConstants_Aes256Gcm),
		string(CryptographySuiteAuthenticationTransformConstants_Md596),
		string(CryptographySuiteAuthenticationTransformConstants_Sha196),
		string(CryptographySuiteAuthenticationTransformConstants_Sha256128),
	}
}

func (s *CryptographySuiteAuthenticationTransformConstants) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCryptographySuiteAuthenticationTransformConstants(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCryptographySuiteAuthenticationTransformConstants(input string) (*CryptographySuiteAuthenticationTransformConstants, error) {
	vals := map[string]CryptographySuiteAuthenticationTransformConstants{
		"aes128gcm":   CryptographySuiteAuthenticationTransformConstants_Aes128Gcm,
		"aes192gcm":   CryptographySuiteAuthenticationTransformConstants_Aes192Gcm,
		"aes256gcm":   CryptographySuiteAuthenticationTransformConstants_Aes256Gcm,
		"md5_96":      CryptographySuiteAuthenticationTransformConstants_Md596,
		"sha1_96":     CryptographySuiteAuthenticationTransformConstants_Sha196,
		"sha_256_128": CryptographySuiteAuthenticationTransformConstants_Sha256128,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteAuthenticationTransformConstants(input)
	return &out, nil
}
