package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerFixedDrivePolicyEncryptionMethod string

const (
	BitLockerFixedDrivePolicyEncryptionMethod_AesCbc128 BitLockerFixedDrivePolicyEncryptionMethod = "aesCbc128"
	BitLockerFixedDrivePolicyEncryptionMethod_AesCbc256 BitLockerFixedDrivePolicyEncryptionMethod = "aesCbc256"
	BitLockerFixedDrivePolicyEncryptionMethod_XtsAes128 BitLockerFixedDrivePolicyEncryptionMethod = "xtsAes128"
	BitLockerFixedDrivePolicyEncryptionMethod_XtsAes256 BitLockerFixedDrivePolicyEncryptionMethod = "xtsAes256"
)

func PossibleValuesForBitLockerFixedDrivePolicyEncryptionMethod() []string {
	return []string{
		string(BitLockerFixedDrivePolicyEncryptionMethod_AesCbc128),
		string(BitLockerFixedDrivePolicyEncryptionMethod_AesCbc256),
		string(BitLockerFixedDrivePolicyEncryptionMethod_XtsAes128),
		string(BitLockerFixedDrivePolicyEncryptionMethod_XtsAes256),
	}
}

func (s *BitLockerFixedDrivePolicyEncryptionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerFixedDrivePolicyEncryptionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerFixedDrivePolicyEncryptionMethod(input string) (*BitLockerFixedDrivePolicyEncryptionMethod, error) {
	vals := map[string]BitLockerFixedDrivePolicyEncryptionMethod{
		"aescbc128": BitLockerFixedDrivePolicyEncryptionMethod_AesCbc128,
		"aescbc256": BitLockerFixedDrivePolicyEncryptionMethod_AesCbc256,
		"xtsaes128": BitLockerFixedDrivePolicyEncryptionMethod_XtsAes128,
		"xtsaes256": BitLockerFixedDrivePolicyEncryptionMethod_XtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerFixedDrivePolicyEncryptionMethod(input)
	return &out, nil
}
