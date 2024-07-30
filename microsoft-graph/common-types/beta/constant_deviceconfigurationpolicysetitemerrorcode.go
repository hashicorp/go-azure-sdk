package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationPolicySetItemErrorCode string

const (
	DeviceConfigurationPolicySetItemErrorCode_Deleted      DeviceConfigurationPolicySetItemErrorCode = "deleted"
	DeviceConfigurationPolicySetItemErrorCode_NoError      DeviceConfigurationPolicySetItemErrorCode = "noError"
	DeviceConfigurationPolicySetItemErrorCode_NotFound     DeviceConfigurationPolicySetItemErrorCode = "notFound"
	DeviceConfigurationPolicySetItemErrorCode_Unauthorized DeviceConfigurationPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForDeviceConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceConfigurationPolicySetItemErrorCode_Deleted),
		string(DeviceConfigurationPolicySetItemErrorCode_NoError),
		string(DeviceConfigurationPolicySetItemErrorCode_NotFound),
		string(DeviceConfigurationPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *DeviceConfigurationPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationPolicySetItemErrorCode(input string) (*DeviceConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]DeviceConfigurationPolicySetItemErrorCode{
		"deleted":      DeviceConfigurationPolicySetItemErrorCode_Deleted,
		"noerror":      DeviceConfigurationPolicySetItemErrorCode_NoError,
		"notfound":     DeviceConfigurationPolicySetItemErrorCode_NotFound,
		"unauthorized": DeviceConfigurationPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
