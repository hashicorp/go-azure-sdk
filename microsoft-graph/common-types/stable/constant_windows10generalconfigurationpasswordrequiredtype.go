package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPasswordRequiredType string

const (
	Windows10GeneralConfigurationPasswordRequiredType_Alphanumeric  Windows10GeneralConfigurationPasswordRequiredType = "alphanumeric"
	Windows10GeneralConfigurationPasswordRequiredType_DeviceDefault Windows10GeneralConfigurationPasswordRequiredType = "deviceDefault"
	Windows10GeneralConfigurationPasswordRequiredType_Numeric       Windows10GeneralConfigurationPasswordRequiredType = "numeric"
)

func PossibleValuesForWindows10GeneralConfigurationPasswordRequiredType() []string {
	return []string{
		string(Windows10GeneralConfigurationPasswordRequiredType_Alphanumeric),
		string(Windows10GeneralConfigurationPasswordRequiredType_DeviceDefault),
		string(Windows10GeneralConfigurationPasswordRequiredType_Numeric),
	}
}

func (s *Windows10GeneralConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPasswordRequiredType(input string) (*Windows10GeneralConfigurationPasswordRequiredType, error) {
	vals := map[string]Windows10GeneralConfigurationPasswordRequiredType{
		"alphanumeric":  Windows10GeneralConfigurationPasswordRequiredType_Alphanumeric,
		"devicedefault": Windows10GeneralConfigurationPasswordRequiredType_DeviceDefault,
		"numeric":       Windows10GeneralConfigurationPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPasswordRequiredType(input)
	return &out, nil
}
