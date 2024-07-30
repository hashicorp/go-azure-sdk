package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCompliancePolicyPasswordRequiredType string

const (
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_Alphabetic              AospDeviceOwnerCompliancePolicyPasswordRequiredType = "alphabetic"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_Alphanumeric            AospDeviceOwnerCompliancePolicyPasswordRequiredType = "alphanumeric"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols AospDeviceOwnerCompliancePolicyPasswordRequiredType = "alphanumericWithSymbols"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_CustomPassword          AospDeviceOwnerCompliancePolicyPasswordRequiredType = "customPassword"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_DeviceDefault           AospDeviceOwnerCompliancePolicyPasswordRequiredType = "deviceDefault"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_LowSecurityBiometric    AospDeviceOwnerCompliancePolicyPasswordRequiredType = "lowSecurityBiometric"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_Numeric                 AospDeviceOwnerCompliancePolicyPasswordRequiredType = "numeric"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_NumericComplex          AospDeviceOwnerCompliancePolicyPasswordRequiredType = "numericComplex"
	AospDeviceOwnerCompliancePolicyPasswordRequiredType_Required                AospDeviceOwnerCompliancePolicyPasswordRequiredType = "required"
)

func PossibleValuesForAospDeviceOwnerCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_Alphabetic),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_CustomPassword),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_LowSecurityBiometric),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_Numeric),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_NumericComplex),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredType_Required),
	}
}

func (s *AospDeviceOwnerCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerCompliancePolicyPasswordRequiredType(input string) (*AospDeviceOwnerCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AospDeviceOwnerCompliancePolicyPasswordRequiredType{
		"alphabetic":              AospDeviceOwnerCompliancePolicyPasswordRequiredType_Alphabetic,
		"alphanumeric":            AospDeviceOwnerCompliancePolicyPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AospDeviceOwnerCompliancePolicyPasswordRequiredType_AlphanumericWithSymbols,
		"custompassword":          AospDeviceOwnerCompliancePolicyPasswordRequiredType_CustomPassword,
		"devicedefault":           AospDeviceOwnerCompliancePolicyPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AospDeviceOwnerCompliancePolicyPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AospDeviceOwnerCompliancePolicyPasswordRequiredType_Numeric,
		"numericcomplex":          AospDeviceOwnerCompliancePolicyPasswordRequiredType_NumericComplex,
		"required":                AospDeviceOwnerCompliancePolicyPasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
