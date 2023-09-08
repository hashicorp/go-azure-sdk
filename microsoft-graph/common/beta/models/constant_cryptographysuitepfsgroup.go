package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuitePfsGroup string

const (
	CryptographySuitePfsGroupecp256  CryptographySuitePfsGroup = "Ecp256"
	CryptographySuitePfsGroupecp384  CryptographySuitePfsGroup = "Ecp384"
	CryptographySuitePfsGrouppfs1    CryptographySuitePfsGroup = "Pfs1"
	CryptographySuitePfsGrouppfs2    CryptographySuitePfsGroup = "Pfs2"
	CryptographySuitePfsGrouppfs2048 CryptographySuitePfsGroup = "Pfs2048"
	CryptographySuitePfsGrouppfs24   CryptographySuitePfsGroup = "Pfs24"
	CryptographySuitePfsGrouppfsMM   CryptographySuitePfsGroup = "PfsMM"
)

func PossibleValuesForCryptographySuitePfsGroup() []string {
	return []string{
		string(CryptographySuitePfsGroupecp256),
		string(CryptographySuitePfsGroupecp384),
		string(CryptographySuitePfsGrouppfs1),
		string(CryptographySuitePfsGrouppfs2),
		string(CryptographySuitePfsGrouppfs2048),
		string(CryptographySuitePfsGrouppfs24),
		string(CryptographySuitePfsGrouppfsMM),
	}
}

func parseCryptographySuitePfsGroup(input string) (*CryptographySuitePfsGroup, error) {
	vals := map[string]CryptographySuitePfsGroup{
		"ecp256":  CryptographySuitePfsGroupecp256,
		"ecp384":  CryptographySuitePfsGroupecp384,
		"pfs1":    CryptographySuitePfsGrouppfs1,
		"pfs2":    CryptographySuitePfsGrouppfs2,
		"pfs2048": CryptographySuitePfsGrouppfs2048,
		"pfs24":   CryptographySuitePfsGrouppfs24,
		"pfsmm":   CryptographySuitePfsGrouppfsMM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuitePfsGroup(input)
	return &out, nil
}
