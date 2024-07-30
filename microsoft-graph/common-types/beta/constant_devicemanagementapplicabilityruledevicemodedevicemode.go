package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleDeviceModeDeviceMode string

const (
	DeviceManagementApplicabilityRuleDeviceModeDeviceMode_SModeConfiguration    DeviceManagementApplicabilityRuleDeviceModeDeviceMode = "sModeConfiguration"
	DeviceManagementApplicabilityRuleDeviceModeDeviceMode_StandardConfiguration DeviceManagementApplicabilityRuleDeviceModeDeviceMode = "standardConfiguration"
)

func PossibleValuesForDeviceManagementApplicabilityRuleDeviceModeDeviceMode() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleDeviceModeDeviceMode_SModeConfiguration),
		string(DeviceManagementApplicabilityRuleDeviceModeDeviceMode_StandardConfiguration),
	}
}

func (s *DeviceManagementApplicabilityRuleDeviceModeDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleDeviceModeDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleDeviceModeDeviceMode(input string) (*DeviceManagementApplicabilityRuleDeviceModeDeviceMode, error) {
	vals := map[string]DeviceManagementApplicabilityRuleDeviceModeDeviceMode{
		"smodeconfiguration":    DeviceManagementApplicabilityRuleDeviceModeDeviceMode_SModeConfiguration,
		"standardconfiguration": DeviceManagementApplicabilityRuleDeviceModeDeviceMode_StandardConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleDeviceModeDeviceMode(input)
	return &out, nil
}
