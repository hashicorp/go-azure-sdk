package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRemovableDrivePolicyEncryptionMethod string

const (
	BitLockerRemovableDrivePolicyEncryptionMethodaesCbc128 BitLockerRemovableDrivePolicyEncryptionMethod = "AesCbc128"
	BitLockerRemovableDrivePolicyEncryptionMethodaesCbc256 BitLockerRemovableDrivePolicyEncryptionMethod = "AesCbc256"
	BitLockerRemovableDrivePolicyEncryptionMethodxtsAes128 BitLockerRemovableDrivePolicyEncryptionMethod = "XtsAes128"
	BitLockerRemovableDrivePolicyEncryptionMethodxtsAes256 BitLockerRemovableDrivePolicyEncryptionMethod = "XtsAes256"
)

func PossibleValuesForBitLockerRemovableDrivePolicyEncryptionMethod() []string {
	return []string{
		string(BitLockerRemovableDrivePolicyEncryptionMethodaesCbc128),
		string(BitLockerRemovableDrivePolicyEncryptionMethodaesCbc256),
		string(BitLockerRemovableDrivePolicyEncryptionMethodxtsAes128),
		string(BitLockerRemovableDrivePolicyEncryptionMethodxtsAes256),
	}
}

func parseBitLockerRemovableDrivePolicyEncryptionMethod(input string) (*BitLockerRemovableDrivePolicyEncryptionMethod, error) {
	vals := map[string]BitLockerRemovableDrivePolicyEncryptionMethod{
		"aescbc128": BitLockerRemovableDrivePolicyEncryptionMethodaesCbc128,
		"aescbc256": BitLockerRemovableDrivePolicyEncryptionMethodaesCbc256,
		"xtsaes128": BitLockerRemovableDrivePolicyEncryptionMethodxtsAes128,
		"xtsaes256": BitLockerRemovableDrivePolicyEncryptionMethodxtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRemovableDrivePolicyEncryptionMethod(input)
	return &out, nil
}
