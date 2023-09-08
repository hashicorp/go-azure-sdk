package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationTemplateType string

const (
	DeviceEnrollmentNotificationConfigurationTemplateTypeemail DeviceEnrollmentNotificationConfigurationTemplateType = "Email"
	DeviceEnrollmentNotificationConfigurationTemplateTypepush  DeviceEnrollmentNotificationConfigurationTemplateType = "Push"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationTemplateType() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationTemplateTypeemail),
		string(DeviceEnrollmentNotificationConfigurationTemplateTypepush),
	}
}

func parseDeviceEnrollmentNotificationConfigurationTemplateType(input string) (*DeviceEnrollmentNotificationConfigurationTemplateType, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationTemplateType{
		"email": DeviceEnrollmentNotificationConfigurationTemplateTypeemail,
		"push":  DeviceEnrollmentNotificationConfigurationTemplateTypepush,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationTemplateType(input)
	return &out, nil
}
