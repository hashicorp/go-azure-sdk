package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Limit                                                 DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Unknown                                               DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
