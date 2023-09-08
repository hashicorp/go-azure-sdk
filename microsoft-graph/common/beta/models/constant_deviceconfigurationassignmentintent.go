package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationAssignmentIntent string

const (
	DeviceConfigurationAssignmentIntentapply  DeviceConfigurationAssignmentIntent = "Apply"
	DeviceConfigurationAssignmentIntentremove DeviceConfigurationAssignmentIntent = "Remove"
)

func PossibleValuesForDeviceConfigurationAssignmentIntent() []string {
	return []string{
		string(DeviceConfigurationAssignmentIntentapply),
		string(DeviceConfigurationAssignmentIntentremove),
	}
}

func parseDeviceConfigurationAssignmentIntent(input string) (*DeviceConfigurationAssignmentIntent, error) {
	vals := map[string]DeviceConfigurationAssignmentIntent{
		"apply":  DeviceConfigurationAssignmentIntentapply,
		"remove": DeviceConfigurationAssignmentIntentremove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationAssignmentIntent(input)
	return &out, nil
}
