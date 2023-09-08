package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertDetectionSource string

const (
	SecurityAlertDetectionSourceantivirus                     SecurityAlertDetectionSource = "Antivirus"
	SecurityAlertDetectionSourceappGovernanceDetection        SecurityAlertDetectionSource = "AppGovernanceDetection"
	SecurityAlertDetectionSourceappGovernancePolicy           SecurityAlertDetectionSource = "AppGovernancePolicy"
	SecurityAlertDetectionSourceautomatedInvestigation        SecurityAlertDetectionSource = "AutomatedInvestigation"
	SecurityAlertDetectionSourceazureAdIdentityProtection     SecurityAlertDetectionSource = "AzureAdIdentityProtection"
	SecurityAlertDetectionSourcecloudAppSecurity              SecurityAlertDetectionSource = "CloudAppSecurity"
	SecurityAlertDetectionSourcecustomDetection               SecurityAlertDetectionSource = "CustomDetection"
	SecurityAlertDetectionSourcecustomTi                      SecurityAlertDetectionSource = "CustomTi"
	SecurityAlertDetectionSourcemanual                        SecurityAlertDetectionSource = "Manual"
	SecurityAlertDetectionSourcemicrosoft365Defender          SecurityAlertDetectionSource = "Microsoft365Defender"
	SecurityAlertDetectionSourcemicrosoftDataLossPrevention   SecurityAlertDetectionSource = "MicrosoftDataLossPrevention"
	SecurityAlertDetectionSourcemicrosoftDefenderForCloud     SecurityAlertDetectionSource = "MicrosoftDefenderForCloud"
	SecurityAlertDetectionSourcemicrosoftDefenderForEndpoint  SecurityAlertDetectionSource = "MicrosoftDefenderForEndpoint"
	SecurityAlertDetectionSourcemicrosoftDefenderForIdentity  SecurityAlertDetectionSource = "MicrosoftDefenderForIdentity"
	SecurityAlertDetectionSourcemicrosoftDefenderForOffice365 SecurityAlertDetectionSource = "MicrosoftDefenderForOffice365"
	SecurityAlertDetectionSourcemicrosoftThreatExperts        SecurityAlertDetectionSource = "MicrosoftThreatExperts"
	SecurityAlertDetectionSourcesmartScreen                   SecurityAlertDetectionSource = "SmartScreen"
	SecurityAlertDetectionSourceunknown                       SecurityAlertDetectionSource = "Unknown"
)

func PossibleValuesForSecurityAlertDetectionSource() []string {
	return []string{
		string(SecurityAlertDetectionSourceantivirus),
		string(SecurityAlertDetectionSourceappGovernanceDetection),
		string(SecurityAlertDetectionSourceappGovernancePolicy),
		string(SecurityAlertDetectionSourceautomatedInvestigation),
		string(SecurityAlertDetectionSourceazureAdIdentityProtection),
		string(SecurityAlertDetectionSourcecloudAppSecurity),
		string(SecurityAlertDetectionSourcecustomDetection),
		string(SecurityAlertDetectionSourcecustomTi),
		string(SecurityAlertDetectionSourcemanual),
		string(SecurityAlertDetectionSourcemicrosoft365Defender),
		string(SecurityAlertDetectionSourcemicrosoftDataLossPrevention),
		string(SecurityAlertDetectionSourcemicrosoftDefenderForCloud),
		string(SecurityAlertDetectionSourcemicrosoftDefenderForEndpoint),
		string(SecurityAlertDetectionSourcemicrosoftDefenderForIdentity),
		string(SecurityAlertDetectionSourcemicrosoftDefenderForOffice365),
		string(SecurityAlertDetectionSourcemicrosoftThreatExperts),
		string(SecurityAlertDetectionSourcesmartScreen),
		string(SecurityAlertDetectionSourceunknown),
	}
}

func parseSecurityAlertDetectionSource(input string) (*SecurityAlertDetectionSource, error) {
	vals := map[string]SecurityAlertDetectionSource{
		"antivirus":                     SecurityAlertDetectionSourceantivirus,
		"appgovernancedetection":        SecurityAlertDetectionSourceappGovernanceDetection,
		"appgovernancepolicy":           SecurityAlertDetectionSourceappGovernancePolicy,
		"automatedinvestigation":        SecurityAlertDetectionSourceautomatedInvestigation,
		"azureadidentityprotection":     SecurityAlertDetectionSourceazureAdIdentityProtection,
		"cloudappsecurity":              SecurityAlertDetectionSourcecloudAppSecurity,
		"customdetection":               SecurityAlertDetectionSourcecustomDetection,
		"customti":                      SecurityAlertDetectionSourcecustomTi,
		"manual":                        SecurityAlertDetectionSourcemanual,
		"microsoft365defender":          SecurityAlertDetectionSourcemicrosoft365Defender,
		"microsoftdatalossprevention":   SecurityAlertDetectionSourcemicrosoftDataLossPrevention,
		"microsoftdefenderforcloud":     SecurityAlertDetectionSourcemicrosoftDefenderForCloud,
		"microsoftdefenderforendpoint":  SecurityAlertDetectionSourcemicrosoftDefenderForEndpoint,
		"microsoftdefenderforidentity":  SecurityAlertDetectionSourcemicrosoftDefenderForIdentity,
		"microsoftdefenderforoffice365": SecurityAlertDetectionSourcemicrosoftDefenderForOffice365,
		"microsoftthreatexperts":        SecurityAlertDetectionSourcemicrosoftThreatExperts,
		"smartscreen":                   SecurityAlertDetectionSourcesmartScreen,
		"unknown":                       SecurityAlertDetectionSourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertDetectionSource(input)
	return &out, nil
}
