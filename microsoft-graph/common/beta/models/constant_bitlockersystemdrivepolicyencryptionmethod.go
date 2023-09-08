package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyEncryptionMethod string

const (
	BitLockerSystemDrivePolicyEncryptionMethodaesCbc128 BitLockerSystemDrivePolicyEncryptionMethod = "AesCbc128"
	BitLockerSystemDrivePolicyEncryptionMethodaesCbc256 BitLockerSystemDrivePolicyEncryptionMethod = "AesCbc256"
	BitLockerSystemDrivePolicyEncryptionMethodxtsAes128 BitLockerSystemDrivePolicyEncryptionMethod = "XtsAes128"
	BitLockerSystemDrivePolicyEncryptionMethodxtsAes256 BitLockerSystemDrivePolicyEncryptionMethod = "XtsAes256"
)

func PossibleValuesForBitLockerSystemDrivePolicyEncryptionMethod() []string {
	return []string{
		string(BitLockerSystemDrivePolicyEncryptionMethodaesCbc128),
		string(BitLockerSystemDrivePolicyEncryptionMethodaesCbc256),
		string(BitLockerSystemDrivePolicyEncryptionMethodxtsAes128),
		string(BitLockerSystemDrivePolicyEncryptionMethodxtsAes256),
	}
}

func parseBitLockerSystemDrivePolicyEncryptionMethod(input string) (*BitLockerSystemDrivePolicyEncryptionMethod, error) {
	vals := map[string]BitLockerSystemDrivePolicyEncryptionMethod{
		"aescbc128": BitLockerSystemDrivePolicyEncryptionMethodaesCbc128,
		"aescbc256": BitLockerSystemDrivePolicyEncryptionMethodaesCbc256,
		"xtsaes128": BitLockerSystemDrivePolicyEncryptionMethodxtsAes128,
		"xtsaes256": BitLockerSystemDrivePolicyEncryptionMethodxtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyEncryptionMethod(input)
	return &out, nil
}
