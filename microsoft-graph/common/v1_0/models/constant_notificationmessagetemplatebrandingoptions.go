package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationMessageTemplateBrandingOptions string

const (
	NotificationMessageTemplateBrandingOptionsincludeCompanyLogo        NotificationMessageTemplateBrandingOptions = "IncludeCompanyLogo"
	NotificationMessageTemplateBrandingOptionsincludeCompanyName        NotificationMessageTemplateBrandingOptions = "IncludeCompanyName"
	NotificationMessageTemplateBrandingOptionsincludeCompanyPortalLink  NotificationMessageTemplateBrandingOptions = "IncludeCompanyPortalLink"
	NotificationMessageTemplateBrandingOptionsincludeContactInformation NotificationMessageTemplateBrandingOptions = "IncludeContactInformation"
	NotificationMessageTemplateBrandingOptionsincludeDeviceDetails      NotificationMessageTemplateBrandingOptions = "IncludeDeviceDetails"
	NotificationMessageTemplateBrandingOptionsnone                      NotificationMessageTemplateBrandingOptions = "None"
)

func PossibleValuesForNotificationMessageTemplateBrandingOptions() []string {
	return []string{
		string(NotificationMessageTemplateBrandingOptionsincludeCompanyLogo),
		string(NotificationMessageTemplateBrandingOptionsincludeCompanyName),
		string(NotificationMessageTemplateBrandingOptionsincludeCompanyPortalLink),
		string(NotificationMessageTemplateBrandingOptionsincludeContactInformation),
		string(NotificationMessageTemplateBrandingOptionsincludeDeviceDetails),
		string(NotificationMessageTemplateBrandingOptionsnone),
	}
}

func parseNotificationMessageTemplateBrandingOptions(input string) (*NotificationMessageTemplateBrandingOptions, error) {
	vals := map[string]NotificationMessageTemplateBrandingOptions{
		"includecompanylogo":        NotificationMessageTemplateBrandingOptionsincludeCompanyLogo,
		"includecompanyname":        NotificationMessageTemplateBrandingOptionsincludeCompanyName,
		"includecompanyportallink":  NotificationMessageTemplateBrandingOptionsincludeCompanyPortalLink,
		"includecontactinformation": NotificationMessageTemplateBrandingOptionsincludeContactInformation,
		"includedevicedetails":      NotificationMessageTemplateBrandingOptionsincludeDeviceDetails,
		"none":                      NotificationMessageTemplateBrandingOptionsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationMessageTemplateBrandingOptions(input)
	return &out, nil
}
