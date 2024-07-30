package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Limit                                                 DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Unknown                                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
