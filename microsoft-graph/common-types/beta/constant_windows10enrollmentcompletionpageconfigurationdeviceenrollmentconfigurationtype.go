package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType string

const (
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Limit                                                 Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "limit"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Unknown                                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "unknown"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForWindows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType(input string) (*Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
