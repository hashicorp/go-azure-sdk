package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyPolicySetItemStatus string

const (
	DeviceManagementConfigurationPolicyPolicySetItemStatus_Error          DeviceManagementConfigurationPolicyPolicySetItemStatus = "error"
	DeviceManagementConfigurationPolicyPolicySetItemStatus_NotAssigned    DeviceManagementConfigurationPolicyPolicySetItemStatus = "notAssigned"
	DeviceManagementConfigurationPolicyPolicySetItemStatus_PartialSuccess DeviceManagementConfigurationPolicyPolicySetItemStatus = "partialSuccess"
	DeviceManagementConfigurationPolicyPolicySetItemStatus_Success        DeviceManagementConfigurationPolicyPolicySetItemStatus = "success"
	DeviceManagementConfigurationPolicyPolicySetItemStatus_Unknown        DeviceManagementConfigurationPolicyPolicySetItemStatus = "unknown"
	DeviceManagementConfigurationPolicyPolicySetItemStatus_Validating     DeviceManagementConfigurationPolicyPolicySetItemStatus = "validating"
)

func PossibleValuesForDeviceManagementConfigurationPolicyPolicySetItemStatus() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyPolicySetItemStatus_Error),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatus_NotAssigned),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatus_PartialSuccess),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatus_Success),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatus_Unknown),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatus_Validating),
	}
}

func (s *DeviceManagementConfigurationPolicyPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyPolicySetItemStatus(input string) (*DeviceManagementConfigurationPolicyPolicySetItemStatus, error) {
	vals := map[string]DeviceManagementConfigurationPolicyPolicySetItemStatus{
		"error":          DeviceManagementConfigurationPolicyPolicySetItemStatus_Error,
		"notassigned":    DeviceManagementConfigurationPolicyPolicySetItemStatus_NotAssigned,
		"partialsuccess": DeviceManagementConfigurationPolicyPolicySetItemStatus_PartialSuccess,
		"success":        DeviceManagementConfigurationPolicyPolicySetItemStatus_Success,
		"unknown":        DeviceManagementConfigurationPolicyPolicySetItemStatus_Unknown,
		"validating":     DeviceManagementConfigurationPolicyPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyPolicySetItemStatus(input)
	return &out, nil
}
