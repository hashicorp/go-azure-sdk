package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationPolicySetItemStatus string

const (
	DeviceConfigurationPolicySetItemStatus_Error          DeviceConfigurationPolicySetItemStatus = "error"
	DeviceConfigurationPolicySetItemStatus_NotAssigned    DeviceConfigurationPolicySetItemStatus = "notAssigned"
	DeviceConfigurationPolicySetItemStatus_PartialSuccess DeviceConfigurationPolicySetItemStatus = "partialSuccess"
	DeviceConfigurationPolicySetItemStatus_Success        DeviceConfigurationPolicySetItemStatus = "success"
	DeviceConfigurationPolicySetItemStatus_Unknown        DeviceConfigurationPolicySetItemStatus = "unknown"
	DeviceConfigurationPolicySetItemStatus_Validating     DeviceConfigurationPolicySetItemStatus = "validating"
)

func PossibleValuesForDeviceConfigurationPolicySetItemStatus() []string {
	return []string{
		string(DeviceConfigurationPolicySetItemStatus_Error),
		string(DeviceConfigurationPolicySetItemStatus_NotAssigned),
		string(DeviceConfigurationPolicySetItemStatus_PartialSuccess),
		string(DeviceConfigurationPolicySetItemStatus_Success),
		string(DeviceConfigurationPolicySetItemStatus_Unknown),
		string(DeviceConfigurationPolicySetItemStatus_Validating),
	}
}

func (s *DeviceConfigurationPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationPolicySetItemStatus(input string) (*DeviceConfigurationPolicySetItemStatus, error) {
	vals := map[string]DeviceConfigurationPolicySetItemStatus{
		"error":          DeviceConfigurationPolicySetItemStatus_Error,
		"notassigned":    DeviceConfigurationPolicySetItemStatus_NotAssigned,
		"partialsuccess": DeviceConfigurationPolicySetItemStatus_PartialSuccess,
		"success":        DeviceConfigurationPolicySetItemStatus_Success,
		"unknown":        DeviceConfigurationPolicySetItemStatus_Unknown,
		"validating":     DeviceConfigurationPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationPolicySetItemStatus(input)
	return &out, nil
}
