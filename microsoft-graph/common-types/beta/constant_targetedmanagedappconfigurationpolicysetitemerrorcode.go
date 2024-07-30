package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationPolicySetItemErrorCode string

const (
	TargetedManagedAppConfigurationPolicySetItemErrorCode_Deleted      TargetedManagedAppConfigurationPolicySetItemErrorCode = "deleted"
	TargetedManagedAppConfigurationPolicySetItemErrorCode_NoError      TargetedManagedAppConfigurationPolicySetItemErrorCode = "noError"
	TargetedManagedAppConfigurationPolicySetItemErrorCode_NotFound     TargetedManagedAppConfigurationPolicySetItemErrorCode = "notFound"
	TargetedManagedAppConfigurationPolicySetItemErrorCode_Unauthorized TargetedManagedAppConfigurationPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForTargetedManagedAppConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(TargetedManagedAppConfigurationPolicySetItemErrorCode_Deleted),
		string(TargetedManagedAppConfigurationPolicySetItemErrorCode_NoError),
		string(TargetedManagedAppConfigurationPolicySetItemErrorCode_NotFound),
		string(TargetedManagedAppConfigurationPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *TargetedManagedAppConfigurationPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppConfigurationPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppConfigurationPolicySetItemErrorCode(input string) (*TargetedManagedAppConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]TargetedManagedAppConfigurationPolicySetItemErrorCode{
		"deleted":      TargetedManagedAppConfigurationPolicySetItemErrorCode_Deleted,
		"noerror":      TargetedManagedAppConfigurationPolicySetItemErrorCode_NoError,
		"notfound":     TargetedManagedAppConfigurationPolicySetItemErrorCode_NotFound,
		"unauthorized": TargetedManagedAppConfigurationPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
