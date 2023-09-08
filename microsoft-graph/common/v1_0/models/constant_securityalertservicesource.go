package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertServiceSource string

const (
	SecurityAlertServiceSourceazureAdIdentityProtection     SecurityAlertServiceSource = "AzureAdIdentityProtection"
	SecurityAlertServiceSourcedataLossPrevention            SecurityAlertServiceSource = "DataLossPrevention"
	SecurityAlertServiceSourcemicrosoft365Defender          SecurityAlertServiceSource = "Microsoft365Defender"
	SecurityAlertServiceSourcemicrosoftAppGovernance        SecurityAlertServiceSource = "MicrosoftAppGovernance"
	SecurityAlertServiceSourcemicrosoftDefenderForCloud     SecurityAlertServiceSource = "MicrosoftDefenderForCloud"
	SecurityAlertServiceSourcemicrosoftDefenderForCloudApps SecurityAlertServiceSource = "MicrosoftDefenderForCloudApps"
	SecurityAlertServiceSourcemicrosoftDefenderForEndpoint  SecurityAlertServiceSource = "MicrosoftDefenderForEndpoint"
	SecurityAlertServiceSourcemicrosoftDefenderForIdentity  SecurityAlertServiceSource = "MicrosoftDefenderForIdentity"
	SecurityAlertServiceSourcemicrosoftDefenderForOffice365 SecurityAlertServiceSource = "MicrosoftDefenderForOffice365"
	SecurityAlertServiceSourceunknown                       SecurityAlertServiceSource = "Unknown"
)

func PossibleValuesForSecurityAlertServiceSource() []string {
	return []string{
		string(SecurityAlertServiceSourceazureAdIdentityProtection),
		string(SecurityAlertServiceSourcedataLossPrevention),
		string(SecurityAlertServiceSourcemicrosoft365Defender),
		string(SecurityAlertServiceSourcemicrosoftAppGovernance),
		string(SecurityAlertServiceSourcemicrosoftDefenderForCloud),
		string(SecurityAlertServiceSourcemicrosoftDefenderForCloudApps),
		string(SecurityAlertServiceSourcemicrosoftDefenderForEndpoint),
		string(SecurityAlertServiceSourcemicrosoftDefenderForIdentity),
		string(SecurityAlertServiceSourcemicrosoftDefenderForOffice365),
		string(SecurityAlertServiceSourceunknown),
	}
}

func parseSecurityAlertServiceSource(input string) (*SecurityAlertServiceSource, error) {
	vals := map[string]SecurityAlertServiceSource{
		"azureadidentityprotection":     SecurityAlertServiceSourceazureAdIdentityProtection,
		"datalossprevention":            SecurityAlertServiceSourcedataLossPrevention,
		"microsoft365defender":          SecurityAlertServiceSourcemicrosoft365Defender,
		"microsoftappgovernance":        SecurityAlertServiceSourcemicrosoftAppGovernance,
		"microsoftdefenderforcloud":     SecurityAlertServiceSourcemicrosoftDefenderForCloud,
		"microsoftdefenderforcloudapps": SecurityAlertServiceSourcemicrosoftDefenderForCloudApps,
		"microsoftdefenderforendpoint":  SecurityAlertServiceSourcemicrosoftDefenderForEndpoint,
		"microsoftdefenderforidentity":  SecurityAlertServiceSourcemicrosoftDefenderForIdentity,
		"microsoftdefenderforoffice365": SecurityAlertServiceSourcemicrosoftDefenderForOffice365,
		"unknown":                       SecurityAlertServiceSourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertServiceSource(input)
	return &out, nil
}
