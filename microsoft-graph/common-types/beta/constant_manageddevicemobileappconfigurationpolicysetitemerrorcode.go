package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode string

const (
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_Deleted      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "deleted"
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_NoError      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "noError"
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_NotFound     ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "notFound"
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_Unauthorized ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_Deleted),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_NoError),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_NotFound),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationPolicySetItemErrorCode(input string) (*ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode{
		"deleted":      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_Deleted,
		"noerror":      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_NoError,
		"notfound":     ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_NotFound,
		"unauthorized": ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
