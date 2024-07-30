package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptPolicySetItemErrorCode string

const (
	DeviceManagementScriptPolicySetItemErrorCode_Deleted      DeviceManagementScriptPolicySetItemErrorCode = "deleted"
	DeviceManagementScriptPolicySetItemErrorCode_NoError      DeviceManagementScriptPolicySetItemErrorCode = "noError"
	DeviceManagementScriptPolicySetItemErrorCode_NotFound     DeviceManagementScriptPolicySetItemErrorCode = "notFound"
	DeviceManagementScriptPolicySetItemErrorCode_Unauthorized DeviceManagementScriptPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForDeviceManagementScriptPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceManagementScriptPolicySetItemErrorCode_Deleted),
		string(DeviceManagementScriptPolicySetItemErrorCode_NoError),
		string(DeviceManagementScriptPolicySetItemErrorCode_NotFound),
		string(DeviceManagementScriptPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *DeviceManagementScriptPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementScriptPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementScriptPolicySetItemErrorCode(input string) (*DeviceManagementScriptPolicySetItemErrorCode, error) {
	vals := map[string]DeviceManagementScriptPolicySetItemErrorCode{
		"deleted":      DeviceManagementScriptPolicySetItemErrorCode_Deleted,
		"noerror":      DeviceManagementScriptPolicySetItemErrorCode_NoError,
		"notfound":     DeviceManagementScriptPolicySetItemErrorCode_NotFound,
		"unauthorized": DeviceManagementScriptPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptPolicySetItemErrorCode(input)
	return &out, nil
}
