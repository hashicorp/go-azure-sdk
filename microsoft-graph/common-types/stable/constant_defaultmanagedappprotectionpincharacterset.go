package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionPinCharacterSet string

const (
	DefaultManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol DefaultManagedAppProtectionPinCharacterSet = "alphanumericAndSymbol"
	DefaultManagedAppProtectionPinCharacterSet_Numeric               DefaultManagedAppProtectionPinCharacterSet = "numeric"
)

func PossibleValuesForDefaultManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(DefaultManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol),
		string(DefaultManagedAppProtectionPinCharacterSet_Numeric),
	}
}

func (s *DefaultManagedAppProtectionPinCharacterSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionPinCharacterSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionPinCharacterSet(input string) (*DefaultManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]DefaultManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": DefaultManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol,
		"numeric":               DefaultManagedAppProtectionPinCharacterSet_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
