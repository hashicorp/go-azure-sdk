package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteIntegrityCheckMethod string

const (
	CryptographySuiteIntegrityCheckMethod_Md5     CryptographySuiteIntegrityCheckMethod = "md5"
	CryptographySuiteIntegrityCheckMethod_Sha1160 CryptographySuiteIntegrityCheckMethod = "sha1_160"
	CryptographySuiteIntegrityCheckMethod_Sha196  CryptographySuiteIntegrityCheckMethod = "sha1_96"
	CryptographySuiteIntegrityCheckMethod_Sha2256 CryptographySuiteIntegrityCheckMethod = "sha2_256"
	CryptographySuiteIntegrityCheckMethod_Sha2384 CryptographySuiteIntegrityCheckMethod = "sha2_384"
	CryptographySuiteIntegrityCheckMethod_Sha2512 CryptographySuiteIntegrityCheckMethod = "sha2_512"
)

func PossibleValuesForCryptographySuiteIntegrityCheckMethod() []string {
	return []string{
		string(CryptographySuiteIntegrityCheckMethod_Md5),
		string(CryptographySuiteIntegrityCheckMethod_Sha1160),
		string(CryptographySuiteIntegrityCheckMethod_Sha196),
		string(CryptographySuiteIntegrityCheckMethod_Sha2256),
		string(CryptographySuiteIntegrityCheckMethod_Sha2384),
		string(CryptographySuiteIntegrityCheckMethod_Sha2512),
	}
}

func (s *CryptographySuiteIntegrityCheckMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCryptographySuiteIntegrityCheckMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCryptographySuiteIntegrityCheckMethod(input string) (*CryptographySuiteIntegrityCheckMethod, error) {
	vals := map[string]CryptographySuiteIntegrityCheckMethod{
		"md5":      CryptographySuiteIntegrityCheckMethod_Md5,
		"sha1_160": CryptographySuiteIntegrityCheckMethod_Sha1160,
		"sha1_96":  CryptographySuiteIntegrityCheckMethod_Sha196,
		"sha2_256": CryptographySuiteIntegrityCheckMethod_Sha2256,
		"sha2_384": CryptographySuiteIntegrityCheckMethod_Sha2384,
		"sha2_512": CryptographySuiteIntegrityCheckMethod_Sha2512,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteIntegrityCheckMethod(input)
	return &out, nil
}
