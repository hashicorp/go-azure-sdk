package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationPasscodeRequiredType string

const (
	IosGeneralDeviceConfigurationPasscodeRequiredTypealphanumeric  IosGeneralDeviceConfigurationPasscodeRequiredType = "Alphanumeric"
	IosGeneralDeviceConfigurationPasscodeRequiredTypedeviceDefault IosGeneralDeviceConfigurationPasscodeRequiredType = "DeviceDefault"
	IosGeneralDeviceConfigurationPasscodeRequiredTypenumeric       IosGeneralDeviceConfigurationPasscodeRequiredType = "Numeric"
)

func PossibleValuesForIosGeneralDeviceConfigurationPasscodeRequiredType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationPasscodeRequiredTypealphanumeric),
		string(IosGeneralDeviceConfigurationPasscodeRequiredTypedeviceDefault),
		string(IosGeneralDeviceConfigurationPasscodeRequiredTypenumeric),
	}
}

func parseIosGeneralDeviceConfigurationPasscodeRequiredType(input string) (*IosGeneralDeviceConfigurationPasscodeRequiredType, error) {
	vals := map[string]IosGeneralDeviceConfigurationPasscodeRequiredType{
		"alphanumeric":  IosGeneralDeviceConfigurationPasscodeRequiredTypealphanumeric,
		"devicedefault": IosGeneralDeviceConfigurationPasscodeRequiredTypedeviceDefault,
		"numeric":       IosGeneralDeviceConfigurationPasscodeRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationPasscodeRequiredType(input)
	return &out, nil
}
