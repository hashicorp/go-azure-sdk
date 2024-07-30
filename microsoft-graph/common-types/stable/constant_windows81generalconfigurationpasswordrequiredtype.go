package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationPasswordRequiredType string

const (
	Windows81GeneralConfigurationPasswordRequiredType_Alphanumeric  Windows81GeneralConfigurationPasswordRequiredType = "alphanumeric"
	Windows81GeneralConfigurationPasswordRequiredType_DeviceDefault Windows81GeneralConfigurationPasswordRequiredType = "deviceDefault"
	Windows81GeneralConfigurationPasswordRequiredType_Numeric       Windows81GeneralConfigurationPasswordRequiredType = "numeric"
)

func PossibleValuesForWindows81GeneralConfigurationPasswordRequiredType() []string {
	return []string{
		string(Windows81GeneralConfigurationPasswordRequiredType_Alphanumeric),
		string(Windows81GeneralConfigurationPasswordRequiredType_DeviceDefault),
		string(Windows81GeneralConfigurationPasswordRequiredType_Numeric),
	}
}

func (s *Windows81GeneralConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationPasswordRequiredType(input string) (*Windows81GeneralConfigurationPasswordRequiredType, error) {
	vals := map[string]Windows81GeneralConfigurationPasswordRequiredType{
		"alphanumeric":  Windows81GeneralConfigurationPasswordRequiredType_Alphanumeric,
		"devicedefault": Windows81GeneralConfigurationPasswordRequiredType_DeviceDefault,
		"numeric":       Windows81GeneralConfigurationPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationPasswordRequiredType(input)
	return &out, nil
}
