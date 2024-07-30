package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicyPasswordRequiredType string

const (
	AndroidForWorkCompliancePolicyPasswordRequiredType_Alphabetic              AndroidForWorkCompliancePolicyPasswordRequiredType = "alphabetic"
	AndroidForWorkCompliancePolicyPasswordRequiredType_Alphanumeric            AndroidForWorkCompliancePolicyPasswordRequiredType = "alphanumeric"
	AndroidForWorkCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols AndroidForWorkCompliancePolicyPasswordRequiredType = "alphanumericWithSymbols"
	AndroidForWorkCompliancePolicyPasswordRequiredType_Any                     AndroidForWorkCompliancePolicyPasswordRequiredType = "any"
	AndroidForWorkCompliancePolicyPasswordRequiredType_DeviceDefault           AndroidForWorkCompliancePolicyPasswordRequiredType = "deviceDefault"
	AndroidForWorkCompliancePolicyPasswordRequiredType_LowSecurityBiometric    AndroidForWorkCompliancePolicyPasswordRequiredType = "lowSecurityBiometric"
	AndroidForWorkCompliancePolicyPasswordRequiredType_Numeric                 AndroidForWorkCompliancePolicyPasswordRequiredType = "numeric"
	AndroidForWorkCompliancePolicyPasswordRequiredType_NumericComplex          AndroidForWorkCompliancePolicyPasswordRequiredType = "numericComplex"
)

func PossibleValuesForAndroidForWorkCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_Alphabetic),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_Any),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_LowSecurityBiometric),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_Numeric),
		string(AndroidForWorkCompliancePolicyPasswordRequiredType_NumericComplex),
	}
}

func (s *AndroidForWorkCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCompliancePolicyPasswordRequiredType(input string) (*AndroidForWorkCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidForWorkCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidForWorkCompliancePolicyPasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidForWorkCompliancePolicyPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidForWorkCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols,
		"any":                     AndroidForWorkCompliancePolicyPasswordRequiredType_Any,
		"devicedefault":           AndroidForWorkCompliancePolicyPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkCompliancePolicyPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidForWorkCompliancePolicyPasswordRequiredType_Numeric,
		"numericcomplex":          AndroidForWorkCompliancePolicyPasswordRequiredType_NumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
