package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers string

const (
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersdisabled      Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersenabled       Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersnotConfigured Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers(input string) (*Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakersnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers(input)
	return &out, nil
}
