package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CompliancePolicyPasswordRequiredType string

const (
	Windows10CompliancePolicyPasswordRequiredType_Alphanumeric  Windows10CompliancePolicyPasswordRequiredType = "alphanumeric"
	Windows10CompliancePolicyPasswordRequiredType_DeviceDefault Windows10CompliancePolicyPasswordRequiredType = "deviceDefault"
	Windows10CompliancePolicyPasswordRequiredType_Numeric       Windows10CompliancePolicyPasswordRequiredType = "numeric"
)

func PossibleValuesForWindows10CompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(Windows10CompliancePolicyPasswordRequiredType_Alphanumeric),
		string(Windows10CompliancePolicyPasswordRequiredType_DeviceDefault),
		string(Windows10CompliancePolicyPasswordRequiredType_Numeric),
	}
}

func (s *Windows10CompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10CompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10CompliancePolicyPasswordRequiredType(input string) (*Windows10CompliancePolicyPasswordRequiredType, error) {
	vals := map[string]Windows10CompliancePolicyPasswordRequiredType{
		"alphanumeric":  Windows10CompliancePolicyPasswordRequiredType_Alphanumeric,
		"devicedefault": Windows10CompliancePolicyPasswordRequiredType_DeviceDefault,
		"numeric":       Windows10CompliancePolicyPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
