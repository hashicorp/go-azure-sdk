package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Alphabetic              AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "alphabetic"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Alphanumeric            AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "alphanumeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "alphanumericWithSymbols"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_CustomPassword          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "customPassword"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault           AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "deviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric    AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "lowSecurityBiometric"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Numeric                 AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "numeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "numericComplex"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required                AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "required"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Alphabetic),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Alphanumeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_CustomPassword),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Numeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType{
		"alphabetic":              AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_CustomPassword,
		"devicedefault":           AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Numeric,
		"numericcomplex":          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex,
		"required":                AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input)
	return &out, nil
}
