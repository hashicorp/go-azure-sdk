package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Allowed    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage = "allowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Disallowed DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage = "disallowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Required   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage = "required"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Allowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Disallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Required),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage{
		"allowed":    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Allowed,
		"disallowed": DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Disallowed,
		"required":   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage(input)
	return &out, nil
}
