package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyEncryptionMethod string

const (
	BitLockerSystemDrivePolicyEncryptionMethod_AesCbc128 BitLockerSystemDrivePolicyEncryptionMethod = "aesCbc128"
	BitLockerSystemDrivePolicyEncryptionMethod_AesCbc256 BitLockerSystemDrivePolicyEncryptionMethod = "aesCbc256"
	BitLockerSystemDrivePolicyEncryptionMethod_XtsAes128 BitLockerSystemDrivePolicyEncryptionMethod = "xtsAes128"
	BitLockerSystemDrivePolicyEncryptionMethod_XtsAes256 BitLockerSystemDrivePolicyEncryptionMethod = "xtsAes256"
)

func PossibleValuesForBitLockerSystemDrivePolicyEncryptionMethod() []string {
	return []string{
		string(BitLockerSystemDrivePolicyEncryptionMethod_AesCbc128),
		string(BitLockerSystemDrivePolicyEncryptionMethod_AesCbc256),
		string(BitLockerSystemDrivePolicyEncryptionMethod_XtsAes128),
		string(BitLockerSystemDrivePolicyEncryptionMethod_XtsAes256),
	}
}

func (s *BitLockerSystemDrivePolicyEncryptionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerSystemDrivePolicyEncryptionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerSystemDrivePolicyEncryptionMethod(input string) (*BitLockerSystemDrivePolicyEncryptionMethod, error) {
	vals := map[string]BitLockerSystemDrivePolicyEncryptionMethod{
		"aescbc128": BitLockerSystemDrivePolicyEncryptionMethod_AesCbc128,
		"aescbc256": BitLockerSystemDrivePolicyEncryptionMethod_AesCbc256,
		"xtsaes128": BitLockerSystemDrivePolicyEncryptionMethod_XtsAes128,
		"xtsaes256": BitLockerSystemDrivePolicyEncryptionMethod_XtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyEncryptionMethod(input)
	return &out, nil
}
