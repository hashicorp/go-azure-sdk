package mhsmprivatelinkresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedHsmSkuFamily string

const (
	ManagedHsmSkuFamilyB ManagedHsmSkuFamily = "B"
	ManagedHsmSkuFamilyC ManagedHsmSkuFamily = "C"
)

func PossibleValuesForManagedHsmSkuFamily() []string {
	return []string{
		string(ManagedHsmSkuFamilyB),
		string(ManagedHsmSkuFamilyC),
	}
}

func (s *ManagedHsmSkuFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedHsmSkuFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedHsmSkuFamily(input string) (*ManagedHsmSkuFamily, error) {
	vals := map[string]ManagedHsmSkuFamily{
		"b": ManagedHsmSkuFamilyB,
		"c": ManagedHsmSkuFamilyC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedHsmSkuFamily(input)
	return &out, nil
}

type ManagedHsmSkuName string

const (
	ManagedHsmSkuNameCustomBSix      ManagedHsmSkuName = "Custom_B6"
	ManagedHsmSkuNameCustomBThreeTwo ManagedHsmSkuName = "Custom_B32"
	ManagedHsmSkuNameCustomCFourTwo  ManagedHsmSkuName = "Custom_C42"
	ManagedHsmSkuNameCustomCOneZero  ManagedHsmSkuName = "Custom_C10"
	ManagedHsmSkuNameStandardBOne    ManagedHsmSkuName = "Standard_B1"
)

func PossibleValuesForManagedHsmSkuName() []string {
	return []string{
		string(ManagedHsmSkuNameCustomBSix),
		string(ManagedHsmSkuNameCustomBThreeTwo),
		string(ManagedHsmSkuNameCustomCFourTwo),
		string(ManagedHsmSkuNameCustomCOneZero),
		string(ManagedHsmSkuNameStandardBOne),
	}
}

func (s *ManagedHsmSkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedHsmSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedHsmSkuName(input string) (*ManagedHsmSkuName, error) {
	vals := map[string]ManagedHsmSkuName{
		"custom_b6":   ManagedHsmSkuNameCustomBSix,
		"custom_b32":  ManagedHsmSkuNameCustomBThreeTwo,
		"custom_c42":  ManagedHsmSkuNameCustomCFourTwo,
		"custom_c10":  ManagedHsmSkuNameCustomCOneZero,
		"standard_b1": ManagedHsmSkuNameStandardBOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedHsmSkuName(input)
	return &out, nil
}
