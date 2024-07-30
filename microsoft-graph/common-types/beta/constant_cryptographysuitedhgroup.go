package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuiteDhGroup string

const (
	CryptographySuiteDhGroup_Ecp256  CryptographySuiteDhGroup = "ecp256"
	CryptographySuiteDhGroup_Ecp384  CryptographySuiteDhGroup = "ecp384"
	CryptographySuiteDhGroup_Group1  CryptographySuiteDhGroup = "group1"
	CryptographySuiteDhGroup_Group14 CryptographySuiteDhGroup = "group14"
	CryptographySuiteDhGroup_Group2  CryptographySuiteDhGroup = "group2"
	CryptographySuiteDhGroup_Group24 CryptographySuiteDhGroup = "group24"
)

func PossibleValuesForCryptographySuiteDhGroup() []string {
	return []string{
		string(CryptographySuiteDhGroup_Ecp256),
		string(CryptographySuiteDhGroup_Ecp384),
		string(CryptographySuiteDhGroup_Group1),
		string(CryptographySuiteDhGroup_Group14),
		string(CryptographySuiteDhGroup_Group2),
		string(CryptographySuiteDhGroup_Group24),
	}
}

func (s *CryptographySuiteDhGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCryptographySuiteDhGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCryptographySuiteDhGroup(input string) (*CryptographySuiteDhGroup, error) {
	vals := map[string]CryptographySuiteDhGroup{
		"ecp256":  CryptographySuiteDhGroup_Ecp256,
		"ecp384":  CryptographySuiteDhGroup_Ecp384,
		"group1":  CryptographySuiteDhGroup_Group1,
		"group14": CryptographySuiteDhGroup_Group14,
		"group2":  CryptographySuiteDhGroup_Group2,
		"group24": CryptographySuiteDhGroup_Group24,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuiteDhGroup(input)
	return &out, nil
}
