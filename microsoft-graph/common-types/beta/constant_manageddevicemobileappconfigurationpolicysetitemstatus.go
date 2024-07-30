package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationPolicySetItemStatus string

const (
	ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Error          ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "error"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatus_NotAssigned    ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "notAssigned"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatus_PartialSuccess ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "partialSuccess"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Success        ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "success"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Unknown        ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "unknown"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Validating     ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "validating"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationPolicySetItemStatus() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Error),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatus_NotAssigned),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatus_PartialSuccess),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Success),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Unknown),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Validating),
	}
}

func (s *ManagedDeviceMobileAppConfigurationPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationPolicySetItemStatus(input string) (*ManagedDeviceMobileAppConfigurationPolicySetItemStatus, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationPolicySetItemStatus{
		"error":          ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Error,
		"notassigned":    ManagedDeviceMobileAppConfigurationPolicySetItemStatus_NotAssigned,
		"partialsuccess": ManagedDeviceMobileAppConfigurationPolicySetItemStatus_PartialSuccess,
		"success":        ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Success,
		"unknown":        ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Unknown,
		"validating":     ManagedDeviceMobileAppConfigurationPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationPolicySetItemStatus(input)
	return &out, nil
}
