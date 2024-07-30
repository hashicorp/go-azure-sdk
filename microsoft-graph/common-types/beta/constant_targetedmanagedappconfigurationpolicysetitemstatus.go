package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationPolicySetItemStatus string

const (
	TargetedManagedAppConfigurationPolicySetItemStatus_Error          TargetedManagedAppConfigurationPolicySetItemStatus = "error"
	TargetedManagedAppConfigurationPolicySetItemStatus_NotAssigned    TargetedManagedAppConfigurationPolicySetItemStatus = "notAssigned"
	TargetedManagedAppConfigurationPolicySetItemStatus_PartialSuccess TargetedManagedAppConfigurationPolicySetItemStatus = "partialSuccess"
	TargetedManagedAppConfigurationPolicySetItemStatus_Success        TargetedManagedAppConfigurationPolicySetItemStatus = "success"
	TargetedManagedAppConfigurationPolicySetItemStatus_Unknown        TargetedManagedAppConfigurationPolicySetItemStatus = "unknown"
	TargetedManagedAppConfigurationPolicySetItemStatus_Validating     TargetedManagedAppConfigurationPolicySetItemStatus = "validating"
)

func PossibleValuesForTargetedManagedAppConfigurationPolicySetItemStatus() []string {
	return []string{
		string(TargetedManagedAppConfigurationPolicySetItemStatus_Error),
		string(TargetedManagedAppConfigurationPolicySetItemStatus_NotAssigned),
		string(TargetedManagedAppConfigurationPolicySetItemStatus_PartialSuccess),
		string(TargetedManagedAppConfigurationPolicySetItemStatus_Success),
		string(TargetedManagedAppConfigurationPolicySetItemStatus_Unknown),
		string(TargetedManagedAppConfigurationPolicySetItemStatus_Validating),
	}
}

func (s *TargetedManagedAppConfigurationPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppConfigurationPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppConfigurationPolicySetItemStatus(input string) (*TargetedManagedAppConfigurationPolicySetItemStatus, error) {
	vals := map[string]TargetedManagedAppConfigurationPolicySetItemStatus{
		"error":          TargetedManagedAppConfigurationPolicySetItemStatus_Error,
		"notassigned":    TargetedManagedAppConfigurationPolicySetItemStatus_NotAssigned,
		"partialsuccess": TargetedManagedAppConfigurationPolicySetItemStatus_PartialSuccess,
		"success":        TargetedManagedAppConfigurationPolicySetItemStatus_Success,
		"unknown":        TargetedManagedAppConfigurationPolicySetItemStatus_Unknown,
		"validating":     TargetedManagedAppConfigurationPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationPolicySetItemStatus(input)
	return &out, nil
}
