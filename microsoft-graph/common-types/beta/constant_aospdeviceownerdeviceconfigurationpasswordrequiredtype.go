package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerDeviceConfigurationPasswordRequiredType string

const (
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Alphabetic              AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "alphabetic"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Alphanumeric            AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "alphanumeric"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "alphanumericWithSymbols"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_CustomPassword          AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "customPassword"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_DeviceDefault           AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "deviceDefault"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_LowSecurityBiometric    AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "lowSecurityBiometric"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Numeric                 AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "numeric"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_NumericComplex          AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "numericComplex"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Required                AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "required"
)

func PossibleValuesForAospDeviceOwnerDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Alphabetic),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Alphanumeric),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_CustomPassword),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_DeviceDefault),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_LowSecurityBiometric),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Numeric),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_NumericComplex),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Required),
	}
}

func (s *AospDeviceOwnerDeviceConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerDeviceConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerDeviceConfigurationPasswordRequiredType(input string) (*AospDeviceOwnerDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AospDeviceOwnerDeviceConfigurationPasswordRequiredType{
		"alphabetic":              AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Alphabetic,
		"alphanumeric":            AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AospDeviceOwnerDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols,
		"custompassword":          AospDeviceOwnerDeviceConfigurationPasswordRequiredType_CustomPassword,
		"devicedefault":           AospDeviceOwnerDeviceConfigurationPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AospDeviceOwnerDeviceConfigurationPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Numeric,
		"numericcomplex":          AospDeviceOwnerDeviceConfigurationPasswordRequiredType_NumericComplex,
		"required":                AospDeviceOwnerDeviceConfigurationPasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
