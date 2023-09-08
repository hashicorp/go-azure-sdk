package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsageallowed    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage = "Allowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsagedisallowed DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage = "Disallowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsagerequired   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage = "Required"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsageallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsagedisallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsagerequired),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage{
		"allowed":    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsageallowed,
		"disallowed": DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsagedisallowed,
		"required":   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage(input)
	return &out, nil
}
