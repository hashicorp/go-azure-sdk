package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionPinCharacterSet string

const (
	ManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol ManagedAppProtectionPinCharacterSet = "alphanumericAndSymbol"
	ManagedAppProtectionPinCharacterSet_Numeric               ManagedAppProtectionPinCharacterSet = "numeric"
)

func PossibleValuesForManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(ManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol),
		string(ManagedAppProtectionPinCharacterSet_Numeric),
	}
}

func (s *ManagedAppProtectionPinCharacterSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionPinCharacterSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionPinCharacterSet(input string) (*ManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]ManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": ManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol,
		"numeric":               ManagedAppProtectionPinCharacterSet_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
