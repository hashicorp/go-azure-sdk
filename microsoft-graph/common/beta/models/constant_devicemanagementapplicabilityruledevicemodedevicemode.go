package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleDeviceModeDeviceMode string

const (
	DeviceManagementApplicabilityRuleDeviceModeDeviceModesModeConfiguration    DeviceManagementApplicabilityRuleDeviceModeDeviceMode = "SModeConfiguration"
	DeviceManagementApplicabilityRuleDeviceModeDeviceModestandardConfiguration DeviceManagementApplicabilityRuleDeviceModeDeviceMode = "StandardConfiguration"
)

func PossibleValuesForDeviceManagementApplicabilityRuleDeviceModeDeviceMode() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleDeviceModeDeviceModesModeConfiguration),
		string(DeviceManagementApplicabilityRuleDeviceModeDeviceModestandardConfiguration),
	}
}

func parseDeviceManagementApplicabilityRuleDeviceModeDeviceMode(input string) (*DeviceManagementApplicabilityRuleDeviceModeDeviceMode, error) {
	vals := map[string]DeviceManagementApplicabilityRuleDeviceModeDeviceMode{
		"smodeconfiguration":    DeviceManagementApplicabilityRuleDeviceModeDeviceModesModeConfiguration,
		"standardconfiguration": DeviceManagementApplicabilityRuleDeviceModeDeviceModestandardConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleDeviceModeDeviceMode(input)
	return &out, nil
}
