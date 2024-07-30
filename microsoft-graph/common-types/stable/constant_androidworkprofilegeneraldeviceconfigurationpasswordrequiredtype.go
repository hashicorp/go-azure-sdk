package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "alphanumericWithSymbols"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphabetic       AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "atLeastAlphabetic"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphanumeric     AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "atLeastAlphanumeric"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastNumeric          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "atLeastNumeric"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault           AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "deviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric    AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "lowSecurityBiometric"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_NumericComplex          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "numericComplex"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_Required                AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "required"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphabetic),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphanumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastNumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_NumericComplex),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_Required),
	}
}

func (s *AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType{
		"alphanumericwithsymbols": AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AlphanumericWithSymbols,
		"atleastalphabetic":       AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphabetic,
		"atleastalphanumeric":     AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastAlphanumeric,
		"atleastnumeric":          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_AtLeastNumeric,
		"devicedefault":           AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_LowSecurityBiometric,
		"numericcomplex":          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_NumericComplex,
		"required":                AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
