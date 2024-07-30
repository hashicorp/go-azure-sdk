package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10MobileCompliancePolicyPasswordRequiredType string

const (
	Windows10MobileCompliancePolicyPasswordRequiredType_Alphanumeric  Windows10MobileCompliancePolicyPasswordRequiredType = "alphanumeric"
	Windows10MobileCompliancePolicyPasswordRequiredType_DeviceDefault Windows10MobileCompliancePolicyPasswordRequiredType = "deviceDefault"
	Windows10MobileCompliancePolicyPasswordRequiredType_Numeric       Windows10MobileCompliancePolicyPasswordRequiredType = "numeric"
)

func PossibleValuesForWindows10MobileCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(Windows10MobileCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(Windows10MobileCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(Windows10MobileCompliancePolicyPasswordRequiredType_Numeric),
	}
}

func (s *Windows10MobileCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10MobileCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10MobileCompliancePolicyPasswordRequiredType(input string) (*Windows10MobileCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]Windows10MobileCompliancePolicyPasswordRequiredType{
		"alphanumeric":  Windows10MobileCompliancePolicyPasswordRequiredType_Alphanumeric,
		"devicedefault": Windows10MobileCompliancePolicyPasswordRequiredType_DeviceDefault,
		"numeric":       Windows10MobileCompliancePolicyPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10MobileCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
