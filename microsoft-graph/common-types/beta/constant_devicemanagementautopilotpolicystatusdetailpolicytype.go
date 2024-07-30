package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyStatusDetailPolicyType string

const (
	DeviceManagementAutopilotPolicyStatusDetailPolicyType_AppModel            DeviceManagementAutopilotPolicyStatusDetailPolicyType = "appModel"
	DeviceManagementAutopilotPolicyStatusDetailPolicyType_Application         DeviceManagementAutopilotPolicyStatusDetailPolicyType = "application"
	DeviceManagementAutopilotPolicyStatusDetailPolicyType_ConfigurationPolicy DeviceManagementAutopilotPolicyStatusDetailPolicyType = "configurationPolicy"
	DeviceManagementAutopilotPolicyStatusDetailPolicyType_Unknown             DeviceManagementAutopilotPolicyStatusDetailPolicyType = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotPolicyStatusDetailPolicyType() []string {
	return []string{
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyType_AppModel),
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyType_Application),
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyType_ConfigurationPolicy),
		string(DeviceManagementAutopilotPolicyStatusDetailPolicyType_Unknown),
	}
}

func (s *DeviceManagementAutopilotPolicyStatusDetailPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotPolicyStatusDetailPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotPolicyStatusDetailPolicyType(input string) (*DeviceManagementAutopilotPolicyStatusDetailPolicyType, error) {
	vals := map[string]DeviceManagementAutopilotPolicyStatusDetailPolicyType{
		"appmodel":            DeviceManagementAutopilotPolicyStatusDetailPolicyType_AppModel,
		"application":         DeviceManagementAutopilotPolicyStatusDetailPolicyType_Application,
		"configurationpolicy": DeviceManagementAutopilotPolicyStatusDetailPolicyType_ConfigurationPolicy,
		"unknown":             DeviceManagementAutopilotPolicyStatusDetailPolicyType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotPolicyStatusDetailPolicyType(input)
	return &out, nil
}
