package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceSdCard string

const (
	Windows10DeviceFirmwareConfigurationInterfaceSdCard_Disabled      Windows10DeviceFirmwareConfigurationInterfaceSdCard = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceSdCard_Enabled       Windows10DeviceFirmwareConfigurationInterfaceSdCard = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceSdCard_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceSdCard = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceSdCard() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceSdCard_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSdCard_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSdCard_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceSdCard) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceSdCard(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceSdCard(input string) (*Windows10DeviceFirmwareConfigurationInterfaceSdCard, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceSdCard{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceSdCard_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceSdCard_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceSdCard_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceSdCard(input)
	return &out, nil
}
