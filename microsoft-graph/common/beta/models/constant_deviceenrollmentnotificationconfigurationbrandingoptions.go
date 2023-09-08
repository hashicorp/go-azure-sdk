package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationBrandingOptions string

const (
	DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyLogo        DeviceEnrollmentNotificationConfigurationBrandingOptions = "IncludeCompanyLogo"
	DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyName        DeviceEnrollmentNotificationConfigurationBrandingOptions = "IncludeCompanyName"
	DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyPortalLink  DeviceEnrollmentNotificationConfigurationBrandingOptions = "IncludeCompanyPortalLink"
	DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeContactInformation DeviceEnrollmentNotificationConfigurationBrandingOptions = "IncludeContactInformation"
	DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeDeviceDetails      DeviceEnrollmentNotificationConfigurationBrandingOptions = "IncludeDeviceDetails"
	DeviceEnrollmentNotificationConfigurationBrandingOptionsnone                      DeviceEnrollmentNotificationConfigurationBrandingOptions = "None"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationBrandingOptions() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyLogo),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyName),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyPortalLink),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeContactInformation),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeDeviceDetails),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptionsnone),
	}
}

func parseDeviceEnrollmentNotificationConfigurationBrandingOptions(input string) (*DeviceEnrollmentNotificationConfigurationBrandingOptions, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationBrandingOptions{
		"includecompanylogo":        DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyLogo,
		"includecompanyname":        DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyName,
		"includecompanyportallink":  DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeCompanyPortalLink,
		"includecontactinformation": DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeContactInformation,
		"includedevicedetails":      DeviceEnrollmentNotificationConfigurationBrandingOptionsincludeDeviceDetails,
		"none":                      DeviceEnrollmentNotificationConfigurationBrandingOptionsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationBrandingOptions(input)
	return &out, nil
}
