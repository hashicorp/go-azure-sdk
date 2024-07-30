package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "alphanumericWithSymbols"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphabetic       AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "atLeastAlphabetic"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphanumeric     AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "atLeastAlphanumeric"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastNumeric          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "atLeastNumeric"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault           AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "deviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric    AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "lowSecurityBiometric"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "numericComplex"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required                AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "required"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphabetic),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphanumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastNumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required),
	}
}

func (s *AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType{
		"alphanumericwithsymbols": AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AlphanumericWithSymbols,
		"atleastalphabetic":       AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphabetic,
		"atleastalphanumeric":     AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastAlphanumeric,
		"atleastnumeric":          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_AtLeastNumeric,
		"devicedefault":           AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_DeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_LowSecurityBiometric,
		"numericcomplex":          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_NumericComplex,
		"required":                AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input)
	return &out, nil
}
