package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionPinCharacterSet string

const (
	TargetedManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol TargetedManagedAppProtectionPinCharacterSet = "alphanumericAndSymbol"
	TargetedManagedAppProtectionPinCharacterSet_Numeric               TargetedManagedAppProtectionPinCharacterSet = "numeric"
)

func PossibleValuesForTargetedManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(TargetedManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol),
		string(TargetedManagedAppProtectionPinCharacterSet_Numeric),
	}
}

func (s *TargetedManagedAppProtectionPinCharacterSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionPinCharacterSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionPinCharacterSet(input string) (*TargetedManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]TargetedManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": TargetedManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol,
		"numeric":               TargetedManagedAppProtectionPinCharacterSet_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
