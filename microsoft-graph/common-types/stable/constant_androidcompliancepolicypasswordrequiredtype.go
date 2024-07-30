package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyPasswordRequiredType string

const (
	AndroidCompliancePolicyPasswordRequiredType_Alphabetic              AndroidCompliancePolicyPasswordRequiredType = "alphabetic"
	AndroidCompliancePolicyPasswordRequiredType_Alphanumeric            AndroidCompliancePolicyPasswordRequiredType = "alphanumeric"
	AndroidCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols AndroidCompliancePolicyPasswordRequiredType = "alphanumericWithSymbols"
	AndroidCompliancePolicyPasswordRequiredType_Any                     AndroidCompliancePolicyPasswordRequiredType = "any"
	AndroidCompliancePolicyPasswordRequiredType_DeviceDefault           AndroidCompliancePolicyPasswordRequiredType = "deviceDefault"
	AndroidCompliancePolicyPasswordRequiredType_LowSecurityBiometric    AndroidCompliancePolicyPasswordRequiredType = "lowSecurityBiometric"
	AndroidCompliancePolicyPasswordRequiredType_Numeric                 AndroidCompliancePolicyPasswordRequiredType = "numeric"
	AndroidCompliancePolicyPasswordRequiredType_NumericComplex          AndroidCompliancePolicyPasswordRequiredType = "numericComplex"
)

func PossibleValuesForAndroidCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidCompliancePolicyPasswordRequiredType_Alphabetic),
		string(AndroidCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(AndroidCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidCompliancePolicyPasswordRequiredType_Any),
		string(AndroidCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(AndroidCompliancePolicyPasswordRequiredType_LowSecurityBiometric),
		string(AndroidCompliancePolicyPasswordRequiredType_Numeric),
		string(AndroidCompliancePolicyPasswordRequiredType_NumericComplex),
	}
}

func (s *AndroidCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCompliancePolicyPasswordRequiredType(input string) (*AndroidCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidCompliancePolicyPasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidCompliancePolicyPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols,
		"any":                     AndroidCompliancePolicyPasswordRequiredType_Any,
		"devicedefault":           AndroidCompliancePolicyPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidCompliancePolicyPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidCompliancePolicyPasswordRequiredType_Numeric,
		"numericcomplex":          AndroidCompliancePolicyPasswordRequiredType_NumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
