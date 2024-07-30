package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyPolicySetItemErrorCode string

const (
	DeviceManagementConfigurationPolicyPolicySetItemErrorCode_Deleted      DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "deleted"
	DeviceManagementConfigurationPolicyPolicySetItemErrorCode_NoError      DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "noError"
	DeviceManagementConfigurationPolicyPolicySetItemErrorCode_NotFound     DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "notFound"
	DeviceManagementConfigurationPolicyPolicySetItemErrorCode_Unauthorized DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForDeviceManagementConfigurationPolicyPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCode_Deleted),
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCode_NoError),
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCode_NotFound),
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *DeviceManagementConfigurationPolicyPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyPolicySetItemErrorCode(input string) (*DeviceManagementConfigurationPolicyPolicySetItemErrorCode, error) {
	vals := map[string]DeviceManagementConfigurationPolicyPolicySetItemErrorCode{
		"deleted":      DeviceManagementConfigurationPolicyPolicySetItemErrorCode_Deleted,
		"noerror":      DeviceManagementConfigurationPolicyPolicySetItemErrorCode_NoError,
		"notfound":     DeviceManagementConfigurationPolicyPolicySetItemErrorCode_NotFound,
		"unauthorized": DeviceManagementConfigurationPolicyPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyPolicySetItemErrorCode(input)
	return &out, nil
}
