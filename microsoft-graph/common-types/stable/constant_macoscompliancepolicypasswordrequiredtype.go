package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyPasswordRequiredType string

const (
	MacOSCompliancePolicyPasswordRequiredType_Alphanumeric  MacOSCompliancePolicyPasswordRequiredType = "alphanumeric"
	MacOSCompliancePolicyPasswordRequiredType_DeviceDefault MacOSCompliancePolicyPasswordRequiredType = "deviceDefault"
	MacOSCompliancePolicyPasswordRequiredType_Numeric       MacOSCompliancePolicyPasswordRequiredType = "numeric"
)

func PossibleValuesForMacOSCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(MacOSCompliancePolicyPasswordRequiredType_Alphanumeric),
		string(MacOSCompliancePolicyPasswordRequiredType_DeviceDefault),
		string(MacOSCompliancePolicyPasswordRequiredType_Numeric),
	}
}

func (s *MacOSCompliancePolicyPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCompliancePolicyPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCompliancePolicyPasswordRequiredType(input string) (*MacOSCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]MacOSCompliancePolicyPasswordRequiredType{
		"alphanumeric":  MacOSCompliancePolicyPasswordRequiredType_Alphanumeric,
		"devicedefault": MacOSCompliancePolicyPasswordRequiredType_DeviceDefault,
		"numeric":       MacOSCompliancePolicyPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
