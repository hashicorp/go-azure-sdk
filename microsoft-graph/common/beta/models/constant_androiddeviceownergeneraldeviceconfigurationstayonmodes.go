package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesac            AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "Ac"
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesnotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesusb           AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "Usb"
	AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModeswireless      AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes = "Wireless"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesac),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesusb),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModeswireless),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes{
		"ac":            AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesac,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesnotConfigured,
		"usb":           AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModesusb,
		"wireless":      AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModeswireless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationStayOnModes(input)
	return &out, nil
}
