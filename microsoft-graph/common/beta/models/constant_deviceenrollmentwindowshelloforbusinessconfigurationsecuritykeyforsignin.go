package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIndisabled      DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn = "Disabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignInenabled       DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn = "Enabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignInnotConfigured DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn = "NotConfigured"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIndisabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignInenabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignInnotConfigured),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn{
		"disabled":      DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIndisabled,
		"enabled":       DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignInenabled,
		"notconfigured": DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignInnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn(input)
	return &out, nil
}
