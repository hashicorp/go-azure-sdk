package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeDeviceClassType string

const (
	DeviceManagementExchangeDeviceClassTypefamily DeviceManagementExchangeDeviceClassType = "Family"
	DeviceManagementExchangeDeviceClassTypemodel  DeviceManagementExchangeDeviceClassType = "Model"
)

func PossibleValuesForDeviceManagementExchangeDeviceClassType() []string {
	return []string{
		string(DeviceManagementExchangeDeviceClassTypefamily),
		string(DeviceManagementExchangeDeviceClassTypemodel),
	}
}

func parseDeviceManagementExchangeDeviceClassType(input string) (*DeviceManagementExchangeDeviceClassType, error) {
	vals := map[string]DeviceManagementExchangeDeviceClassType{
		"family": DeviceManagementExchangeDeviceClassTypefamily,
		"model":  DeviceManagementExchangeDeviceClassTypemodel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeDeviceClassType(input)
	return &out, nil
}
