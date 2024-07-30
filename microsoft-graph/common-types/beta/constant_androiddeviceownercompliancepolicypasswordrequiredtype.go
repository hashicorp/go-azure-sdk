package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicyPasswordRequiredType string

const (
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Alphabetic              AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "alphabetic"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Alphanumeric            AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "alphanumeric"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "alphanumericWithSymbols"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_CustomPassword          AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "customPassword"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_DeviceDefault           AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "deviceDefault"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_LowSecurityBiometric    AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "lowSecurityBiometric"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Numeric                 AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "numeric"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_NumericComplex          AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "numericComplex"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Required                AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "required"
)

func PossibleValuesForAndroidDeviceOwnerCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Alphabetic),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_CustomPassword),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_LowSecurityBiometric),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Numeric),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_NumericComplex),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Required),
	}
}

func (s *AndroidDeviceOwnerCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCompliancePolicyPasswordRequiredType(input string) (*AndroidDeviceOwnerCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidDeviceOwnerCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_CustomPassword,
		"devicedefault":           AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Numeric,
		"numericcomplex":          AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_NumericComplex,
		"required":                AndroidDeviceOwnerCompliancePolicyPasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
