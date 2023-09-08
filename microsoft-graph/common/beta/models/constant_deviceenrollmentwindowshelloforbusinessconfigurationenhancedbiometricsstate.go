package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStatedisabled      DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState = "Disabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStateenabled       DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState = "Enabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStatenotConfigured DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState = "NotConfigured"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStatedisabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStateenabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStatenotConfigured),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState{
		"disabled":      DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStatedisabled,
		"enabled":       DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStateenabled,
		"notconfigured": DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsStatenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState(input)
	return &out, nil
}
