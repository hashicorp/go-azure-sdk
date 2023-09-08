package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsageallowed    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage = "Allowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsagedisallowed DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage = "Disallowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsagerequired   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage = "Required"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsageallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsagedisallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsagerequired),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage{
		"allowed":    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsageallowed,
		"disallowed": DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsagedisallowed,
		"required":   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage(input)
	return &out, nil
}
