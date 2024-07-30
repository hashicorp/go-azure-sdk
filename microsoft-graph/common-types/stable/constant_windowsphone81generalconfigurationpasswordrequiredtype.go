package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81GeneralConfigurationPasswordRequiredType string

const (
	WindowsPhone81GeneralConfigurationPasswordRequiredType_Alphanumeric  WindowsPhone81GeneralConfigurationPasswordRequiredType = "alphanumeric"
	WindowsPhone81GeneralConfigurationPasswordRequiredType_DeviceDefault WindowsPhone81GeneralConfigurationPasswordRequiredType = "deviceDefault"
	WindowsPhone81GeneralConfigurationPasswordRequiredType_Numeric       WindowsPhone81GeneralConfigurationPasswordRequiredType = "numeric"
)

func PossibleValuesForWindowsPhone81GeneralConfigurationPasswordRequiredType() []string {
	return []string{
		string(WindowsPhone81GeneralConfigurationPasswordRequiredType_Alphanumeric),
		string(WindowsPhone81GeneralConfigurationPasswordRequiredType_DeviceDefault),
		string(WindowsPhone81GeneralConfigurationPasswordRequiredType_Numeric),
	}
}

func (s *WindowsPhone81GeneralConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81GeneralConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81GeneralConfigurationPasswordRequiredType(input string) (*WindowsPhone81GeneralConfigurationPasswordRequiredType, error) {
	vals := map[string]WindowsPhone81GeneralConfigurationPasswordRequiredType{
		"alphanumeric":  WindowsPhone81GeneralConfigurationPasswordRequiredType_Alphanumeric,
		"devicedefault": WindowsPhone81GeneralConfigurationPasswordRequiredType_DeviceDefault,
		"numeric":       WindowsPhone81GeneralConfigurationPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81GeneralConfigurationPasswordRequiredType(input)
	return &out, nil
}
