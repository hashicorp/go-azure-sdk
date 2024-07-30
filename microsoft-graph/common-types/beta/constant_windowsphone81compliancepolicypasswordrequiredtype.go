package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CompliancePolicyPasswordRequiredType string

const (
	WindowsPhone81CompliancePolicyPasswordRequiredType_Alphanumeric  WindowsPhone81CompliancePolicyPasswordRequiredType = "alphanumeric"
	WindowsPhone81CompliancePolicyPasswordRequiredType_DeviceDefault WindowsPhone81CompliancePolicyPasswordRequiredType = "deviceDefault"
	WindowsPhone81CompliancePolicyPasswordRequiredType_Numeric       WindowsPhone81CompliancePolicyPasswordRequiredType = "numeric"
)

func PossibleValuesForWindowsPhone81CompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(WindowsPhone81CompliancePolicyPasswordRequiredType_Alphanumeric),
		string(WindowsPhone81CompliancePolicyPasswordRequiredType_DeviceDefault),
		string(WindowsPhone81CompliancePolicyPasswordRequiredType_Numeric),
	}
}

func (s *WindowsPhone81CompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81CompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81CompliancePolicyPasswordRequiredType(input string) (*WindowsPhone81CompliancePolicyPasswordRequiredType, error) {
	vals := map[string]WindowsPhone81CompliancePolicyPasswordRequiredType{
		"alphanumeric":  WindowsPhone81CompliancePolicyPasswordRequiredType_Alphanumeric,
		"devicedefault": WindowsPhone81CompliancePolicyPasswordRequiredType_DeviceDefault,
		"numeric":       WindowsPhone81CompliancePolicyPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
