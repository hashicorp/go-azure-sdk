package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "alphanumericWithSymbols"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphabetic       AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "atLeastAlphabetic"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphanumeric     AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "atLeastAlphanumeric"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastNumeric          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "atLeastNumeric"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault           AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "deviceDefault"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric    AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "lowSecurityBiometric"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "numericComplex"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required                AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "required"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphabetic),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphanumeric),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastNumeric),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType{
		"alphanumericwithsymbols": AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols,
		"atleastalphabetic":       AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphabetic,
		"atleastalphanumeric":     AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphanumeric,
		"atleastnumeric":          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastNumeric,
		"devicedefault":           AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric,
		"numericcomplex":          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex,
		"required":                AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input)
	return &out, nil
}
