package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Alphabetic              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "alphabetic"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric            AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "alphanumeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "alphanumericWithSymbols"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_CustomPassword          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "customPassword"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "deviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "lowSecurityBiometric"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Numeric                 AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "numeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_NumericComplex          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "numericComplex"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Required                AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "required"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Alphabetic),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_CustomPassword),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Numeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_NumericComplex),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Required),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType{
		"alphabetic":              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_CustomPassword,
		"devicedefault":           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Numeric,
		"numericcomplex":          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_NumericComplex,
		"required":                AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
