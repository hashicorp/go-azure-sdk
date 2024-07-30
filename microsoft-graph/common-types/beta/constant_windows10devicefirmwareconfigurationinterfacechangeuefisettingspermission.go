package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission string

const (
	Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission_None              Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission = "none"
	Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission_NotConfiguredOnly Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission = "notConfiguredOnly"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission_None),
		string(Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission_NotConfiguredOnly),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission(input string) (*Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission{
		"none":              Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission_None,
		"notconfiguredonly": Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission_NotConfiguredOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission(input)
	return &out, nil
}
