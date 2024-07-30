package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyPasswordRequiredType string

const (
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_Alphabetic              AndroidWorkProfileCompliancePolicyPasswordRequiredType = "alphabetic"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_Alphanumeric            AndroidWorkProfileCompliancePolicyPasswordRequiredType = "alphanumeric"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols AndroidWorkProfileCompliancePolicyPasswordRequiredType = "alphanumericWithSymbols"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_Any                     AndroidWorkProfileCompliancePolicyPasswordRequiredType = "any"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_DeviceDefault           AndroidWorkProfileCompliancePolicyPasswordRequiredType = "deviceDefault"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_LowSecurityBiometric    AndroidWorkProfileCompliancePolicyPasswordRequiredType = "lowSecurityBiometric"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_Numeric                 AndroidWorkProfileCompliancePolicyPasswordRequiredType = "numeric"
	AndroidWorkProfileCompliancePolicyPasswordRequiredType_NumericComplex          AndroidWorkProfileCompliancePolicyPasswordRequiredType = "numericComplex"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_Alphabetic),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_Any),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_LowSecurityBiometric),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_Numeric),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredType_NumericComplex),
	}
}

func (s *AndroidWorkProfileCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCompliancePolicyPasswordRequiredType(input string) (*AndroidWorkProfileCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidWorkProfileCompliancePolicyPasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidWorkProfileCompliancePolicyPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidWorkProfileCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols,
		"any":                     AndroidWorkProfileCompliancePolicyPasswordRequiredType_Any,
		"devicedefault":           AndroidWorkProfileCompliancePolicyPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileCompliancePolicyPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidWorkProfileCompliancePolicyPasswordRequiredType_Numeric,
		"numericcomplex":          AndroidWorkProfileCompliancePolicyPasswordRequiredType_NumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
