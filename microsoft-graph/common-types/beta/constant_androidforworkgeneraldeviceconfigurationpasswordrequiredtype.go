package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "alphanumericWithSymbols"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphabetic       AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "atLeastAlphabetic"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphanumeric     AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "atLeastAlphanumeric"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastNumeric          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "atLeastNumeric"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault           AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "deviceDefault"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric    AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "lowSecurityBiometric"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_NumericComplex          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "numericComplex"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_Required                AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "required"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphabetic),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphanumeric),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastNumeric),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_NumericComplex),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_Required),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType{
		"alphanumericwithsymbols": AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols,
		"atleastalphabetic":       AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphabetic,
		"atleastalphanumeric":     AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphanumeric,
		"atleastnumeric":          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_AtLeastNumeric,
		"devicedefault":           AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric,
		"numericcomplex":          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_NumericComplex,
		"required":                AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
