package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Allowed    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage = "allowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Disallowed DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage = "disallowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Required   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage = "required"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Allowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Disallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Required),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage{
		"allowed":    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Allowed,
		"disallowed": DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Disallowed,
		"required":   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage(input)
	return &out, nil
}
