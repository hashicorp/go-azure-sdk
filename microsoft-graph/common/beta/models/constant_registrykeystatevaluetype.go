package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyStateValueType string

const (
	RegistryKeyStateValueTypebinary            RegistryKeyStateValueType = "Binary"
	RegistryKeyStateValueTypedword             RegistryKeyStateValueType = "Dword"
	RegistryKeyStateValueTypedwordBigEndian    RegistryKeyStateValueType = "DwordBigEndian"
	RegistryKeyStateValueTypedwordLittleEndian RegistryKeyStateValueType = "DwordLittleEndian"
	RegistryKeyStateValueTypeexpandSz          RegistryKeyStateValueType = "ExpandSz"
	RegistryKeyStateValueTypelink              RegistryKeyStateValueType = "Link"
	RegistryKeyStateValueTypemultiSz           RegistryKeyStateValueType = "MultiSz"
	RegistryKeyStateValueTypenone              RegistryKeyStateValueType = "None"
	RegistryKeyStateValueTypeqword             RegistryKeyStateValueType = "Qword"
	RegistryKeyStateValueTypeqwordlittleEndian RegistryKeyStateValueType = "QwordlittleEndian"
	RegistryKeyStateValueTypesz                RegistryKeyStateValueType = "Sz"
	RegistryKeyStateValueTypeunknown           RegistryKeyStateValueType = "Unknown"
)

func PossibleValuesForRegistryKeyStateValueType() []string {
	return []string{
		string(RegistryKeyStateValueTypebinary),
		string(RegistryKeyStateValueTypedword),
		string(RegistryKeyStateValueTypedwordBigEndian),
		string(RegistryKeyStateValueTypedwordLittleEndian),
		string(RegistryKeyStateValueTypeexpandSz),
		string(RegistryKeyStateValueTypelink),
		string(RegistryKeyStateValueTypemultiSz),
		string(RegistryKeyStateValueTypenone),
		string(RegistryKeyStateValueTypeqword),
		string(RegistryKeyStateValueTypeqwordlittleEndian),
		string(RegistryKeyStateValueTypesz),
		string(RegistryKeyStateValueTypeunknown),
	}
}

func parseRegistryKeyStateValueType(input string) (*RegistryKeyStateValueType, error) {
	vals := map[string]RegistryKeyStateValueType{
		"binary":            RegistryKeyStateValueTypebinary,
		"dword":             RegistryKeyStateValueTypedword,
		"dwordbigendian":    RegistryKeyStateValueTypedwordBigEndian,
		"dwordlittleendian": RegistryKeyStateValueTypedwordLittleEndian,
		"expandsz":          RegistryKeyStateValueTypeexpandSz,
		"link":              RegistryKeyStateValueTypelink,
		"multisz":           RegistryKeyStateValueTypemultiSz,
		"none":              RegistryKeyStateValueTypenone,
		"qword":             RegistryKeyStateValueTypeqword,
		"qwordlittleendian": RegistryKeyStateValueTypeqwordlittleEndian,
		"sz":                RegistryKeyStateValueTypesz,
		"unknown":           RegistryKeyStateValueTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryKeyStateValueType(input)
	return &out, nil
}
