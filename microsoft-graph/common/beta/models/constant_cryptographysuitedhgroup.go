package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteDhGroup string

const (
	CryptographySuiteDhGroupecp256  CryptographySuiteDhGroup = "Ecp256"
	CryptographySuiteDhGroupecp384  CryptographySuiteDhGroup = "Ecp384"
	CryptographySuiteDhGroupgroup1  CryptographySuiteDhGroup = "Group1"
	CryptographySuiteDhGroupgroup14 CryptographySuiteDhGroup = "Group14"
	CryptographySuiteDhGroupgroup2  CryptographySuiteDhGroup = "Group2"
	CryptographySuiteDhGroupgroup24 CryptographySuiteDhGroup = "Group24"
)

func PossibleValuesForCryptographySuiteDhGroup() []string {
	return []string{
		string(CryptographySuiteDhGroupecp256),
		string(CryptographySuiteDhGroupecp384),
		string(CryptographySuiteDhGroupgroup1),
		string(CryptographySuiteDhGroupgroup14),
		string(CryptographySuiteDhGroupgroup2),
		string(CryptographySuiteDhGroupgroup24),
	}
}

func parseCryptographySuiteDhGroup(input string) (*CryptographySuiteDhGroup, error) {
	vals := map[string]CryptographySuiteDhGroup{
		"ecp256":  CryptographySuiteDhGroupecp256,
		"ecp384":  CryptographySuiteDhGroupecp384,
		"group1":  CryptographySuiteDhGroupgroup1,
		"group14": CryptographySuiteDhGroupgroup14,
		"group2":  CryptographySuiteDhGroupgroup2,
		"group24": CryptographySuiteDhGroupgroup24,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteDhGroup(input)
	return &out, nil
}
