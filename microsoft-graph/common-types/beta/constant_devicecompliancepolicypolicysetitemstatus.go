package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyPolicySetItemStatus string

const (
	DeviceCompliancePolicyPolicySetItemStatus_Error          DeviceCompliancePolicyPolicySetItemStatus = "error"
	DeviceCompliancePolicyPolicySetItemStatus_NotAssigned    DeviceCompliancePolicyPolicySetItemStatus = "notAssigned"
	DeviceCompliancePolicyPolicySetItemStatus_PartialSuccess DeviceCompliancePolicyPolicySetItemStatus = "partialSuccess"
	DeviceCompliancePolicyPolicySetItemStatus_Success        DeviceCompliancePolicyPolicySetItemStatus = "success"
	DeviceCompliancePolicyPolicySetItemStatus_Unknown        DeviceCompliancePolicyPolicySetItemStatus = "unknown"
	DeviceCompliancePolicyPolicySetItemStatus_Validating     DeviceCompliancePolicyPolicySetItemStatus = "validating"
)

func PossibleValuesForDeviceCompliancePolicyPolicySetItemStatus() []string {
	return []string{
		string(DeviceCompliancePolicyPolicySetItemStatus_Error),
		string(DeviceCompliancePolicyPolicySetItemStatus_NotAssigned),
		string(DeviceCompliancePolicyPolicySetItemStatus_PartialSuccess),
		string(DeviceCompliancePolicyPolicySetItemStatus_Success),
		string(DeviceCompliancePolicyPolicySetItemStatus_Unknown),
		string(DeviceCompliancePolicyPolicySetItemStatus_Validating),
	}
}

func (s *DeviceCompliancePolicyPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyPolicySetItemStatus(input string) (*DeviceCompliancePolicyPolicySetItemStatus, error) {
	vals := map[string]DeviceCompliancePolicyPolicySetItemStatus{
		"error":          DeviceCompliancePolicyPolicySetItemStatus_Error,
		"notassigned":    DeviceCompliancePolicyPolicySetItemStatus_NotAssigned,
		"partialsuccess": DeviceCompliancePolicyPolicySetItemStatus_PartialSuccess,
		"success":        DeviceCompliancePolicyPolicySetItemStatus_Success,
		"unknown":        DeviceCompliancePolicyPolicySetItemStatus_Unknown,
		"validating":     DeviceCompliancePolicyPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyPolicySetItemStatus(input)
	return &out, nil
}
