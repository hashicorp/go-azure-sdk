package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialSettingsIssuer string

const (
	DeviceManagementDerivedCredentialSettingsIssuerentrustDatacard DeviceManagementDerivedCredentialSettingsIssuer = "EntrustDatacard"
	DeviceManagementDerivedCredentialSettingsIssuerintercede       DeviceManagementDerivedCredentialSettingsIssuer = "Intercede"
	DeviceManagementDerivedCredentialSettingsIssuerpurebred        DeviceManagementDerivedCredentialSettingsIssuer = "Purebred"
	DeviceManagementDerivedCredentialSettingsIssuerxTec            DeviceManagementDerivedCredentialSettingsIssuer = "XTec"
)

func PossibleValuesForDeviceManagementDerivedCredentialSettingsIssuer() []string {
	return []string{
		string(DeviceManagementDerivedCredentialSettingsIssuerentrustDatacard),
		string(DeviceManagementDerivedCredentialSettingsIssuerintercede),
		string(DeviceManagementDerivedCredentialSettingsIssuerpurebred),
		string(DeviceManagementDerivedCredentialSettingsIssuerxTec),
	}
}

func parseDeviceManagementDerivedCredentialSettingsIssuer(input string) (*DeviceManagementDerivedCredentialSettingsIssuer, error) {
	vals := map[string]DeviceManagementDerivedCredentialSettingsIssuer{
		"entrustdatacard": DeviceManagementDerivedCredentialSettingsIssuerentrustDatacard,
		"intercede":       DeviceManagementDerivedCredentialSettingsIssuerintercede,
		"purebred":        DeviceManagementDerivedCredentialSettingsIssuerpurebred,
		"xtec":            DeviceManagementDerivedCredentialSettingsIssuerxTec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDerivedCredentialSettingsIssuer(input)
	return &out, nil
}
