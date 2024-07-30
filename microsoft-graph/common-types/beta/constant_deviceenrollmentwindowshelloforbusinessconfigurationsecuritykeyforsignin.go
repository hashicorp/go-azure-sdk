package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_Disabled      DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn = "disabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_Enabled       DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn = "enabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_NotConfigured DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn = "notConfigured"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_Disabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_Enabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_NotConfigured),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn{
		"disabled":      DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_Disabled,
		"enabled":       DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_Enabled,
		"notconfigured": DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn(input)
	return &out, nil
}
