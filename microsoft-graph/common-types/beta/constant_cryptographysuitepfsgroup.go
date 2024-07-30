package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuitePfsGroup string

const (
	CryptographySuitePfsGroup_Ecp256  CryptographySuitePfsGroup = "ecp256"
	CryptographySuitePfsGroup_Ecp384  CryptographySuitePfsGroup = "ecp384"
	CryptographySuitePfsGroup_Pfs1    CryptographySuitePfsGroup = "pfs1"
	CryptographySuitePfsGroup_Pfs2    CryptographySuitePfsGroup = "pfs2"
	CryptographySuitePfsGroup_Pfs2048 CryptographySuitePfsGroup = "pfs2048"
	CryptographySuitePfsGroup_Pfs24   CryptographySuitePfsGroup = "pfs24"
	CryptographySuitePfsGroup_PfsMM   CryptographySuitePfsGroup = "pfsMM"
)

func PossibleValuesForCryptographySuitePfsGroup() []string {
	return []string{
		string(CryptographySuitePfsGroup_Ecp256),
		string(CryptographySuitePfsGroup_Ecp384),
		string(CryptographySuitePfsGroup_Pfs1),
		string(CryptographySuitePfsGroup_Pfs2),
		string(CryptographySuitePfsGroup_Pfs2048),
		string(CryptographySuitePfsGroup_Pfs24),
		string(CryptographySuitePfsGroup_PfsMM),
	}
}

func (s *CryptographySuitePfsGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCryptographySuitePfsGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCryptographySuitePfsGroup(input string) (*CryptographySuitePfsGroup, error) {
	vals := map[string]CryptographySuitePfsGroup{
		"ecp256":  CryptographySuitePfsGroup_Ecp256,
		"ecp384":  CryptographySuitePfsGroup_Ecp384,
		"pfs1":    CryptographySuitePfsGroup_Pfs1,
		"pfs2":    CryptographySuitePfsGroup_Pfs2,
		"pfs2048": CryptographySuitePfsGroup_Pfs2048,
		"pfs24":   CryptographySuitePfsGroup_Pfs24,
		"pfsmm":   CryptographySuitePfsGroup_PfsMM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CryptographySuitePfsGroup(input)
	return &out, nil
}
