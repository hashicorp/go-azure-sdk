package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationState string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationStatedisabled      DeviceEnrollmentWindowsHelloForBusinessConfigurationState = "Disabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationStateenabled       DeviceEnrollmentWindowsHelloForBusinessConfigurationState = "Enabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationStatenotConfigured DeviceEnrollmentWindowsHelloForBusinessConfigurationState = "NotConfigured"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationState() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationStatedisabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationStateenabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationStatenotConfigured),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationState(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationState, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationState{
		"disabled":      DeviceEnrollmentWindowsHelloForBusinessConfigurationStatedisabled,
		"enabled":       DeviceEnrollmentWindowsHelloForBusinessConfigurationStateenabled,
		"notconfigured": DeviceEnrollmentWindowsHelloForBusinessConfigurationStatenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationState(input)
	return &out, nil
}
