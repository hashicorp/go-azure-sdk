package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyStateValueType string

const (
	RegistryKeyStateValueType_Binary            RegistryKeyStateValueType = "binary"
	RegistryKeyStateValueType_Dword             RegistryKeyStateValueType = "dword"
	RegistryKeyStateValueType_DwordBigEndian    RegistryKeyStateValueType = "dwordBigEndian"
	RegistryKeyStateValueType_DwordLittleEndian RegistryKeyStateValueType = "dwordLittleEndian"
	RegistryKeyStateValueType_ExpandSz          RegistryKeyStateValueType = "expandSz"
	RegistryKeyStateValueType_Link              RegistryKeyStateValueType = "link"
	RegistryKeyStateValueType_MultiSz           RegistryKeyStateValueType = "multiSz"
	RegistryKeyStateValueType_None              RegistryKeyStateValueType = "none"
	RegistryKeyStateValueType_Qword             RegistryKeyStateValueType = "qword"
	RegistryKeyStateValueType_QwordlittleEndian RegistryKeyStateValueType = "qwordlittleEndian"
	RegistryKeyStateValueType_Sz                RegistryKeyStateValueType = "sz"
	RegistryKeyStateValueType_Unknown           RegistryKeyStateValueType = "unknown"
)

func PossibleValuesForRegistryKeyStateValueType() []string {
	return []string{
		string(RegistryKeyStateValueType_Binary),
		string(RegistryKeyStateValueType_Dword),
		string(RegistryKeyStateValueType_DwordBigEndian),
		string(RegistryKeyStateValueType_DwordLittleEndian),
		string(RegistryKeyStateValueType_ExpandSz),
		string(RegistryKeyStateValueType_Link),
		string(RegistryKeyStateValueType_MultiSz),
		string(RegistryKeyStateValueType_None),
		string(RegistryKeyStateValueType_Qword),
		string(RegistryKeyStateValueType_QwordlittleEndian),
		string(RegistryKeyStateValueType_Sz),
		string(RegistryKeyStateValueType_Unknown),
	}
}

func (s *RegistryKeyStateValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryKeyStateValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryKeyStateValueType(input string) (*RegistryKeyStateValueType, error) {
	vals := map[string]RegistryKeyStateValueType{
		"binary":            RegistryKeyStateValueType_Binary,
		"dword":             RegistryKeyStateValueType_Dword,
		"dwordbigendian":    RegistryKeyStateValueType_DwordBigEndian,
		"dwordlittleendian": RegistryKeyStateValueType_DwordLittleEndian,
		"expandsz":          RegistryKeyStateValueType_ExpandSz,
		"link":              RegistryKeyStateValueType_Link,
		"multisz":           RegistryKeyStateValueType_MultiSz,
		"none":              RegistryKeyStateValueType_None,
		"qword":             RegistryKeyStateValueType_Qword,
		"qwordlittleendian": RegistryKeyStateValueType_QwordlittleEndian,
		"sz":                RegistryKeyStateValueType_Sz,
		"unknown":           RegistryKeyStateValueType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryKeyStateValueType(input)
	return &out, nil
}
