package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRemovableDrivePolicyEncryptionMethod string

const (
	BitLockerRemovableDrivePolicyEncryptionMethod_AesCbc128 BitLockerRemovableDrivePolicyEncryptionMethod = "aesCbc128"
	BitLockerRemovableDrivePolicyEncryptionMethod_AesCbc256 BitLockerRemovableDrivePolicyEncryptionMethod = "aesCbc256"
	BitLockerRemovableDrivePolicyEncryptionMethod_XtsAes128 BitLockerRemovableDrivePolicyEncryptionMethod = "xtsAes128"
	BitLockerRemovableDrivePolicyEncryptionMethod_XtsAes256 BitLockerRemovableDrivePolicyEncryptionMethod = "xtsAes256"
)

func PossibleValuesForBitLockerRemovableDrivePolicyEncryptionMethod() []string {
	return []string{
		string(BitLockerRemovableDrivePolicyEncryptionMethod_AesCbc128),
		string(BitLockerRemovableDrivePolicyEncryptionMethod_AesCbc256),
		string(BitLockerRemovableDrivePolicyEncryptionMethod_XtsAes128),
		string(BitLockerRemovableDrivePolicyEncryptionMethod_XtsAes256),
	}
}

func (s *BitLockerRemovableDrivePolicyEncryptionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerRemovableDrivePolicyEncryptionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerRemovableDrivePolicyEncryptionMethod(input string) (*BitLockerRemovableDrivePolicyEncryptionMethod, error) {
	vals := map[string]BitLockerRemovableDrivePolicyEncryptionMethod{
		"aescbc128": BitLockerRemovableDrivePolicyEncryptionMethod_AesCbc128,
		"aescbc256": BitLockerRemovableDrivePolicyEncryptionMethod_AesCbc256,
		"xtsaes128": BitLockerRemovableDrivePolicyEncryptionMethod_XtsAes128,
		"xtsaes256": BitLockerRemovableDrivePolicyEncryptionMethod_XtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRemovableDrivePolicyEncryptionMethod(input)
	return &out, nil
}
