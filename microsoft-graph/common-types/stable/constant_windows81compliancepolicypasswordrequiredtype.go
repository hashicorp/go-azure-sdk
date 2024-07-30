package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CompliancePolicyPasswordRequiredType string

const (
	Windows81CompliancePolicyPasswordRequiredType_Alphanumeric  Windows81CompliancePolicyPasswordRequiredType = "alphanumeric"
	Windows81CompliancePolicyPasswordRequiredType_DeviceDefault Windows81CompliancePolicyPasswordRequiredType = "deviceDefault"
	Windows81CompliancePolicyPasswordRequiredType_Numeric       Windows81CompliancePolicyPasswordRequiredType = "numeric"
)

func PossibleValuesForWindows81CompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(Windows81CompliancePolicyPasswordRequiredType_Alphanumeric),
		string(Windows81CompliancePolicyPasswordRequiredType_DeviceDefault),
		string(Windows81CompliancePolicyPasswordRequiredType_Numeric),
	}
}

func (s *Windows81CompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81CompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81CompliancePolicyPasswordRequiredType(input string) (*Windows81CompliancePolicyPasswordRequiredType, error) {
	vals := map[string]Windows81CompliancePolicyPasswordRequiredType{
		"alphanumeric":  Windows81CompliancePolicyPasswordRequiredType_Alphanumeric,
		"devicedefault": Windows81CompliancePolicyPasswordRequiredType_DeviceDefault,
		"numeric":       Windows81CompliancePolicyPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
