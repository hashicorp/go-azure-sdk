package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerFixedDrivePolicyEncryptionMethod string

const (
	BitLockerFixedDrivePolicyEncryptionMethodaesCbc128 BitLockerFixedDrivePolicyEncryptionMethod = "AesCbc128"
	BitLockerFixedDrivePolicyEncryptionMethodaesCbc256 BitLockerFixedDrivePolicyEncryptionMethod = "AesCbc256"
	BitLockerFixedDrivePolicyEncryptionMethodxtsAes128 BitLockerFixedDrivePolicyEncryptionMethod = "XtsAes128"
	BitLockerFixedDrivePolicyEncryptionMethodxtsAes256 BitLockerFixedDrivePolicyEncryptionMethod = "XtsAes256"
)

func PossibleValuesForBitLockerFixedDrivePolicyEncryptionMethod() []string {
	return []string{
		string(BitLockerFixedDrivePolicyEncryptionMethodaesCbc128),
		string(BitLockerFixedDrivePolicyEncryptionMethodaesCbc256),
		string(BitLockerFixedDrivePolicyEncryptionMethodxtsAes128),
		string(BitLockerFixedDrivePolicyEncryptionMethodxtsAes256),
	}
}

func parseBitLockerFixedDrivePolicyEncryptionMethod(input string) (*BitLockerFixedDrivePolicyEncryptionMethod, error) {
	vals := map[string]BitLockerFixedDrivePolicyEncryptionMethod{
		"aescbc128": BitLockerFixedDrivePolicyEncryptionMethodaesCbc128,
		"aescbc256": BitLockerFixedDrivePolicyEncryptionMethodaesCbc256,
		"xtsaes128": BitLockerFixedDrivePolicyEncryptionMethodxtsAes128,
		"xtsaes256": BitLockerFixedDrivePolicyEncryptionMethodxtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerFixedDrivePolicyEncryptionMethod(input)
	return &out, nil
}
