package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionPinCharacterSet string

const (
	AndroidManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol AndroidManagedAppProtectionPinCharacterSet = "alphanumericAndSymbol"
	AndroidManagedAppProtectionPinCharacterSet_Numeric               AndroidManagedAppProtectionPinCharacterSet = "numeric"
)

func PossibleValuesForAndroidManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(AndroidManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol),
		string(AndroidManagedAppProtectionPinCharacterSet_Numeric),
	}
}

func (s *AndroidManagedAppProtectionPinCharacterSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionPinCharacterSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionPinCharacterSet(input string) (*AndroidManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]AndroidManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": AndroidManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol,
		"numeric":               AndroidManagedAppProtectionPinCharacterSet_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
