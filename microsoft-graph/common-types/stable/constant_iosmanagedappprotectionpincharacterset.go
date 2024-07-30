package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionPinCharacterSet string

const (
	IosManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol IosManagedAppProtectionPinCharacterSet = "alphanumericAndSymbol"
	IosManagedAppProtectionPinCharacterSet_Numeric               IosManagedAppProtectionPinCharacterSet = "numeric"
)

func PossibleValuesForIosManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(IosManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol),
		string(IosManagedAppProtectionPinCharacterSet_Numeric),
	}
}

func (s *IosManagedAppProtectionPinCharacterSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionPinCharacterSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionPinCharacterSet(input string) (*IosManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]IosManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": IosManagedAppProtectionPinCharacterSet_AlphanumericAndSymbol,
		"numeric":               IosManagedAppProtectionPinCharacterSet_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
