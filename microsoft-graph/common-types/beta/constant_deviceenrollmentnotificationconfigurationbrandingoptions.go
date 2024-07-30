package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationBrandingOptions string

const (
	DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyLogo        DeviceEnrollmentNotificationConfigurationBrandingOptions = "includeCompanyLogo"
	DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyName        DeviceEnrollmentNotificationConfigurationBrandingOptions = "includeCompanyName"
	DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyPortalLink  DeviceEnrollmentNotificationConfigurationBrandingOptions = "includeCompanyPortalLink"
	DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeContactInformation DeviceEnrollmentNotificationConfigurationBrandingOptions = "includeContactInformation"
	DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeDeviceDetails      DeviceEnrollmentNotificationConfigurationBrandingOptions = "includeDeviceDetails"
	DeviceEnrollmentNotificationConfigurationBrandingOptions_None                      DeviceEnrollmentNotificationConfigurationBrandingOptions = "none"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationBrandingOptions() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyLogo),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyName),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyPortalLink),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeContactInformation),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeDeviceDetails),
		string(DeviceEnrollmentNotificationConfigurationBrandingOptions_None),
	}
}

func (s *DeviceEnrollmentNotificationConfigurationBrandingOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentNotificationConfigurationBrandingOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentNotificationConfigurationBrandingOptions(input string) (*DeviceEnrollmentNotificationConfigurationBrandingOptions, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationBrandingOptions{
		"includecompanylogo":        DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyLogo,
		"includecompanyname":        DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyName,
		"includecompanyportallink":  DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeCompanyPortalLink,
		"includecontactinformation": DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeContactInformation,
		"includedevicedetails":      DeviceEnrollmentNotificationConfigurationBrandingOptions_IncludeDeviceDetails,
		"none":                      DeviceEnrollmentNotificationConfigurationBrandingOptions_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationBrandingOptions(input)
	return &out, nil
}
