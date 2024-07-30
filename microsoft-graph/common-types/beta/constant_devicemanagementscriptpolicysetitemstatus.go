package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptPolicySetItemStatus string

const (
	DeviceManagementScriptPolicySetItemStatus_Error          DeviceManagementScriptPolicySetItemStatus = "error"
	DeviceManagementScriptPolicySetItemStatus_NotAssigned    DeviceManagementScriptPolicySetItemStatus = "notAssigned"
	DeviceManagementScriptPolicySetItemStatus_PartialSuccess DeviceManagementScriptPolicySetItemStatus = "partialSuccess"
	DeviceManagementScriptPolicySetItemStatus_Success        DeviceManagementScriptPolicySetItemStatus = "success"
	DeviceManagementScriptPolicySetItemStatus_Unknown        DeviceManagementScriptPolicySetItemStatus = "unknown"
	DeviceManagementScriptPolicySetItemStatus_Validating     DeviceManagementScriptPolicySetItemStatus = "validating"
)

func PossibleValuesForDeviceManagementScriptPolicySetItemStatus() []string {
	return []string{
		string(DeviceManagementScriptPolicySetItemStatus_Error),
		string(DeviceManagementScriptPolicySetItemStatus_NotAssigned),
		string(DeviceManagementScriptPolicySetItemStatus_PartialSuccess),
		string(DeviceManagementScriptPolicySetItemStatus_Success),
		string(DeviceManagementScriptPolicySetItemStatus_Unknown),
		string(DeviceManagementScriptPolicySetItemStatus_Validating),
	}
}

func (s *DeviceManagementScriptPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementScriptPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementScriptPolicySetItemStatus(input string) (*DeviceManagementScriptPolicySetItemStatus, error) {
	vals := map[string]DeviceManagementScriptPolicySetItemStatus{
		"error":          DeviceManagementScriptPolicySetItemStatus_Error,
		"notassigned":    DeviceManagementScriptPolicySetItemStatus_NotAssigned,
		"partialsuccess": DeviceManagementScriptPolicySetItemStatus_PartialSuccess,
		"success":        DeviceManagementScriptPolicySetItemStatus_Success,
		"unknown":        DeviceManagementScriptPolicySetItemStatus_Unknown,
		"validating":     DeviceManagementScriptPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptPolicySetItemStatus(input)
	return &out, nil
}
