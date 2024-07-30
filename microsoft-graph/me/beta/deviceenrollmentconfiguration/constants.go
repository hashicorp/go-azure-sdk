package deviceenrollmentconfiguration

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone),
	}
}

func (s *DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude,
		"include": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude,
		"none":    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}

type DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultLimit                                          DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultPlatformRestrictions                           DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindowsHelloForBusiness                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDeviceComanagementAuthorityConfiguration              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeEnrollmentNotificationsConfiguration                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeLimit                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypePlatformRestrictions                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeSinglePlatformRestriction                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeUnknown                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindowsHelloForBusiness                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultLimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultPlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeEnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeLimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypePlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeSinglePlatformRestriction),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeUnknown),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeEnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeLimit,
		"platformrestrictions":                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypePlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeSinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeUnknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}

type EnrollmentConfigurationAssignmentSource string

const (
	EnrollmentConfigurationAssignmentSourceDirect     EnrollmentConfigurationAssignmentSource = "direct"
	EnrollmentConfigurationAssignmentSourcePolicySets EnrollmentConfigurationAssignmentSource = "policySets"
)

func PossibleValuesForEnrollmentConfigurationAssignmentSource() []string {
	return []string{
		string(EnrollmentConfigurationAssignmentSourceDirect),
		string(EnrollmentConfigurationAssignmentSourcePolicySets),
	}
}

func (s *EnrollmentConfigurationAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentConfigurationAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentConfigurationAssignmentSource(input string) (*EnrollmentConfigurationAssignmentSource, error) {
	vals := map[string]EnrollmentConfigurationAssignmentSource{
		"direct":     EnrollmentConfigurationAssignmentSourceDirect,
		"policysets": EnrollmentConfigurationAssignmentSourcePolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentConfigurationAssignmentSource(input)
	return &out, nil
}
