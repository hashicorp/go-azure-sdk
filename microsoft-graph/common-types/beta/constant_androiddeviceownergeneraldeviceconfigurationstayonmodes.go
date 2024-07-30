package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Ac            AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "ac"
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Usb           AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "usb"
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Wireless      AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "wireless"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Ac),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Usb),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Wireless),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes{
		"ac":            AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Ac,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_NotConfigured,
		"usb":           AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Usb,
		"wireless":      AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes_Wireless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes(input)
	return &out, nil
}
