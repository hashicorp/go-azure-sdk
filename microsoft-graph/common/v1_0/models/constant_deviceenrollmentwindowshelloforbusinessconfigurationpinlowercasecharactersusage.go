package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsageallowed    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage = "Allowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsagedisallowed DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage = "Disallowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsagerequired   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage = "Required"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsageallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsagedisallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsagerequired),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage{
		"allowed":    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsageallowed,
		"disallowed": DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsagedisallowed,
		"required":   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage(input)
	return &out, nil
}
