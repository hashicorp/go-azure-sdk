package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationMessageTemplateBrandingOptions string

const (
	NotificationMessageTemplateBrandingOptions_IncludeCompanyLogo        NotificationMessageTemplateBrandingOptions = "includeCompanyLogo"
	NotificationMessageTemplateBrandingOptions_IncludeCompanyName        NotificationMessageTemplateBrandingOptions = "includeCompanyName"
	NotificationMessageTemplateBrandingOptions_IncludeCompanyPortalLink  NotificationMessageTemplateBrandingOptions = "includeCompanyPortalLink"
	NotificationMessageTemplateBrandingOptions_IncludeContactInformation NotificationMessageTemplateBrandingOptions = "includeContactInformation"
	NotificationMessageTemplateBrandingOptions_IncludeDeviceDetails      NotificationMessageTemplateBrandingOptions = "includeDeviceDetails"
	NotificationMessageTemplateBrandingOptions_None                      NotificationMessageTemplateBrandingOptions = "none"
)

func PossibleValuesForNotificationMessageTemplateBrandingOptions() []string {
	return []string{
		string(NotificationMessageTemplateBrandingOptions_IncludeCompanyLogo),
		string(NotificationMessageTemplateBrandingOptions_IncludeCompanyName),
		string(NotificationMessageTemplateBrandingOptions_IncludeCompanyPortalLink),
		string(NotificationMessageTemplateBrandingOptions_IncludeContactInformation),
		string(NotificationMessageTemplateBrandingOptions_IncludeDeviceDetails),
		string(NotificationMessageTemplateBrandingOptions_None),
	}
}

func (s *NotificationMessageTemplateBrandingOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationMessageTemplateBrandingOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationMessageTemplateBrandingOptions(input string) (*NotificationMessageTemplateBrandingOptions, error) {
	vals := map[string]NotificationMessageTemplateBrandingOptions{
		"includecompanylogo":        NotificationMessageTemplateBrandingOptions_IncludeCompanyLogo,
		"includecompanyname":        NotificationMessageTemplateBrandingOptions_IncludeCompanyName,
		"includecompanyportallink":  NotificationMessageTemplateBrandingOptions_IncludeCompanyPortalLink,
		"includecontactinformation": NotificationMessageTemplateBrandingOptions_IncludeContactInformation,
		"includedevicedetails":      NotificationMessageTemplateBrandingOptions_IncludeDeviceDetails,
		"none":                      NotificationMessageTemplateBrandingOptions_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationMessageTemplateBrandingOptions(input)
	return &out, nil
}
