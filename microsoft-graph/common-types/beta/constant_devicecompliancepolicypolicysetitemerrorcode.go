package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyPolicySetItemErrorCode string

const (
	DeviceCompliancePolicyPolicySetItemErrorCode_Deleted      DeviceCompliancePolicyPolicySetItemErrorCode = "deleted"
	DeviceCompliancePolicyPolicySetItemErrorCode_NoError      DeviceCompliancePolicyPolicySetItemErrorCode = "noError"
	DeviceCompliancePolicyPolicySetItemErrorCode_NotFound     DeviceCompliancePolicyPolicySetItemErrorCode = "notFound"
	DeviceCompliancePolicyPolicySetItemErrorCode_Unauthorized DeviceCompliancePolicyPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForDeviceCompliancePolicyPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceCompliancePolicyPolicySetItemErrorCode_Deleted),
		string(DeviceCompliancePolicyPolicySetItemErrorCode_NoError),
		string(DeviceCompliancePolicyPolicySetItemErrorCode_NotFound),
		string(DeviceCompliancePolicyPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *DeviceCompliancePolicyPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyPolicySetItemErrorCode(input string) (*DeviceCompliancePolicyPolicySetItemErrorCode, error) {
	vals := map[string]DeviceCompliancePolicyPolicySetItemErrorCode{
		"deleted":      DeviceCompliancePolicyPolicySetItemErrorCode_Deleted,
		"noerror":      DeviceCompliancePolicyPolicySetItemErrorCode_NoError,
		"notfound":     DeviceCompliancePolicyPolicySetItemErrorCode_NotFound,
		"unauthorized": DeviceCompliancePolicyPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyPolicySetItemErrorCode(input)
	return &out, nil
}
