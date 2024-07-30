package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyPasscodeRequiredType string

const (
	IosCompliancePolicyPasscodeRequiredType_Alphanumeric  IosCompliancePolicyPasscodeRequiredType = "alphanumeric"
	IosCompliancePolicyPasscodeRequiredType_DeviceDefault IosCompliancePolicyPasscodeRequiredType = "deviceDefault"
	IosCompliancePolicyPasscodeRequiredType_Numeric       IosCompliancePolicyPasscodeRequiredType = "numeric"
)

func PossibleValuesForIosCompliancePolicyPasscodeRequiredType() []string {
	return []string{
		string(IosCompliancePolicyPasscodeRequiredType_Alphanumeric),
		string(IosCompliancePolicyPasscodeRequiredType_DeviceDefault),
		string(IosCompliancePolicyPasscodeRequiredType_Numeric),
	}
}

func (s *IosCompliancePolicyPasscodeRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosCompliancePolicyPasscodeRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosCompliancePolicyPasscodeRequiredType(input string) (*IosCompliancePolicyPasscodeRequiredType, error) {
	vals := map[string]IosCompliancePolicyPasscodeRequiredType{
		"alphanumeric":  IosCompliancePolicyPasscodeRequiredType_Alphanumeric,
		"devicedefault": IosCompliancePolicyPasscodeRequiredType_DeviceDefault,
		"numeric":       IosCompliancePolicyPasscodeRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCompliancePolicyPasscodeRequiredType(input)
	return &out, nil
}
