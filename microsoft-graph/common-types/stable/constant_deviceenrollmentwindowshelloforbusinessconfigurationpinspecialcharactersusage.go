package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Allowed    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage = "allowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Disallowed DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage = "disallowed"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Required   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage = "required"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Allowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Disallowed),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Required),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage{
		"allowed":    DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Allowed,
		"disallowed": DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Disallowed,
		"required":   DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage(input)
	return &out, nil
}
