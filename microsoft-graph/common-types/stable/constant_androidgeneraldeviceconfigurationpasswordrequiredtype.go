package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidGeneralDeviceConfigurationPasswordRequiredType_Alphabetic              AndroidGeneralDeviceConfigurationPasswordRequiredType = "alphabetic"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric            AndroidGeneralDeviceConfigurationPasswordRequiredType = "alphanumeric"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols AndroidGeneralDeviceConfigurationPasswordRequiredType = "alphanumericWithSymbols"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_Any                     AndroidGeneralDeviceConfigurationPasswordRequiredType = "any"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault           AndroidGeneralDeviceConfigurationPasswordRequiredType = "deviceDefault"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric    AndroidGeneralDeviceConfigurationPasswordRequiredType = "lowSecurityBiometric"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_Numeric                 AndroidGeneralDeviceConfigurationPasswordRequiredType = "numeric"
	AndroidGeneralDeviceConfigurationPasswordRequiredType_NumericComplex          AndroidGeneralDeviceConfigurationPasswordRequiredType = "numericComplex"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_Alphabetic),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_Any),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_Numeric),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredType_NumericComplex),
	}
}

func (s *AndroidGeneralDeviceConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidGeneralDeviceConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationPasswordRequiredType{
		"alphabetic":              AndroidGeneralDeviceConfigurationPasswordRequiredType_Alphabetic,
		"alphanumeric":            AndroidGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric,
		"alphanumericwithsymbols": AndroidGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols,
		"any":                     AndroidGeneralDeviceConfigurationPasswordRequiredType_Any,
		"devicedefault":           AndroidGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric,
		"numeric":                 AndroidGeneralDeviceConfigurationPasswordRequiredType_Numeric,
		"numericcomplex":          AndroidGeneralDeviceConfigurationPasswordRequiredType_NumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
