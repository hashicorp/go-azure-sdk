package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateLifecycleState string

const (
	DeviceManagementConfigurationPolicyTemplateLifecycleState_Active     DeviceManagementConfigurationPolicyTemplateLifecycleState = "active"
	DeviceManagementConfigurationPolicyTemplateLifecycleState_Deprecated DeviceManagementConfigurationPolicyTemplateLifecycleState = "deprecated"
	DeviceManagementConfigurationPolicyTemplateLifecycleState_Draft      DeviceManagementConfigurationPolicyTemplateLifecycleState = "draft"
	DeviceManagementConfigurationPolicyTemplateLifecycleState_Invalid    DeviceManagementConfigurationPolicyTemplateLifecycleState = "invalid"
	DeviceManagementConfigurationPolicyTemplateLifecycleState_Retired    DeviceManagementConfigurationPolicyTemplateLifecycleState = "retired"
	DeviceManagementConfigurationPolicyTemplateLifecycleState_Superseded DeviceManagementConfigurationPolicyTemplateLifecycleState = "superseded"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateLifecycleState() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateLifecycleState_Active),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleState_Deprecated),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleState_Draft),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleState_Invalid),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleState_Retired),
		string(DeviceManagementConfigurationPolicyTemplateLifecycleState_Superseded),
	}
}

func (s *DeviceManagementConfigurationPolicyTemplateLifecycleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyTemplateLifecycleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyTemplateLifecycleState(input string) (*DeviceManagementConfigurationPolicyTemplateLifecycleState, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateLifecycleState{
		"active":     DeviceManagementConfigurationPolicyTemplateLifecycleState_Active,
		"deprecated": DeviceManagementConfigurationPolicyTemplateLifecycleState_Deprecated,
		"draft":      DeviceManagementConfigurationPolicyTemplateLifecycleState_Draft,
		"invalid":    DeviceManagementConfigurationPolicyTemplateLifecycleState_Invalid,
		"retired":    DeviceManagementConfigurationPolicyTemplateLifecycleState_Retired,
		"superseded": DeviceManagementConfigurationPolicyTemplateLifecycleState_Superseded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateLifecycleState(input)
	return &out, nil
}
