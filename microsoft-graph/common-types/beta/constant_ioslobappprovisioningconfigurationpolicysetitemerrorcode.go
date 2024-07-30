package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppProvisioningConfigurationPolicySetItemErrorCode string

const (
	IosLobAppProvisioningConfigurationPolicySetItemErrorCode_Deleted      IosLobAppProvisioningConfigurationPolicySetItemErrorCode = "deleted"
	IosLobAppProvisioningConfigurationPolicySetItemErrorCode_NoError      IosLobAppProvisioningConfigurationPolicySetItemErrorCode = "noError"
	IosLobAppProvisioningConfigurationPolicySetItemErrorCode_NotFound     IosLobAppProvisioningConfigurationPolicySetItemErrorCode = "notFound"
	IosLobAppProvisioningConfigurationPolicySetItemErrorCode_Unauthorized IosLobAppProvisioningConfigurationPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForIosLobAppProvisioningConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(IosLobAppProvisioningConfigurationPolicySetItemErrorCode_Deleted),
		string(IosLobAppProvisioningConfigurationPolicySetItemErrorCode_NoError),
		string(IosLobAppProvisioningConfigurationPolicySetItemErrorCode_NotFound),
		string(IosLobAppProvisioningConfigurationPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *IosLobAppProvisioningConfigurationPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosLobAppProvisioningConfigurationPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosLobAppProvisioningConfigurationPolicySetItemErrorCode(input string) (*IosLobAppProvisioningConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]IosLobAppProvisioningConfigurationPolicySetItemErrorCode{
		"deleted":      IosLobAppProvisioningConfigurationPolicySetItemErrorCode_Deleted,
		"noerror":      IosLobAppProvisioningConfigurationPolicySetItemErrorCode_NoError,
		"notfound":     IosLobAppProvisioningConfigurationPolicySetItemErrorCode_NotFound,
		"unauthorized": IosLobAppProvisioningConfigurationPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosLobAppProvisioningConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
