package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppProvisioningConfigurationPolicySetItemStatus string

const (
	IosLobAppProvisioningConfigurationPolicySetItemStatus_Error          IosLobAppProvisioningConfigurationPolicySetItemStatus = "error"
	IosLobAppProvisioningConfigurationPolicySetItemStatus_NotAssigned    IosLobAppProvisioningConfigurationPolicySetItemStatus = "notAssigned"
	IosLobAppProvisioningConfigurationPolicySetItemStatus_PartialSuccess IosLobAppProvisioningConfigurationPolicySetItemStatus = "partialSuccess"
	IosLobAppProvisioningConfigurationPolicySetItemStatus_Success        IosLobAppProvisioningConfigurationPolicySetItemStatus = "success"
	IosLobAppProvisioningConfigurationPolicySetItemStatus_Unknown        IosLobAppProvisioningConfigurationPolicySetItemStatus = "unknown"
	IosLobAppProvisioningConfigurationPolicySetItemStatus_Validating     IosLobAppProvisioningConfigurationPolicySetItemStatus = "validating"
)

func PossibleValuesForIosLobAppProvisioningConfigurationPolicySetItemStatus() []string {
	return []string{
		string(IosLobAppProvisioningConfigurationPolicySetItemStatus_Error),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatus_NotAssigned),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatus_PartialSuccess),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatus_Success),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatus_Unknown),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatus_Validating),
	}
}

func (s *IosLobAppProvisioningConfigurationPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosLobAppProvisioningConfigurationPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosLobAppProvisioningConfigurationPolicySetItemStatus(input string) (*IosLobAppProvisioningConfigurationPolicySetItemStatus, error) {
	vals := map[string]IosLobAppProvisioningConfigurationPolicySetItemStatus{
		"error":          IosLobAppProvisioningConfigurationPolicySetItemStatus_Error,
		"notassigned":    IosLobAppProvisioningConfigurationPolicySetItemStatus_NotAssigned,
		"partialsuccess": IosLobAppProvisioningConfigurationPolicySetItemStatus_PartialSuccess,
		"success":        IosLobAppProvisioningConfigurationPolicySetItemStatus_Success,
		"unknown":        IosLobAppProvisioningConfigurationPolicySetItemStatus_Unknown,
		"validating":     IosLobAppProvisioningConfigurationPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosLobAppProvisioningConfigurationPolicySetItemStatus(input)
	return &out, nil
}
