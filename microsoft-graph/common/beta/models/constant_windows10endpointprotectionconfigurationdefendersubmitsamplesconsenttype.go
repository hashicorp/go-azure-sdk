package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType string

const (
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypealwaysPrompt                 Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "AlwaysPrompt"
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypeneverSend                    Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "NeverSend"
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypesendAllSamplesAutomatically  Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "SendAllSamplesAutomatically"
	Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypesendSafeSamplesAutomatically Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType = "SendSafeSamplesAutomatically"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypealwaysPrompt),
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypeneverSend),
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypesendAllSamplesAutomatically),
		string(Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypesendSafeSamplesAutomatically),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType(input string) (*Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType{
		"alwaysprompt":                 Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypealwaysPrompt,
		"neversend":                    Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypeneverSend,
		"sendallsamplesautomatically":  Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypesendAllSamplesAutomatically,
		"sendsafesamplesautomatically": Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentTypesendSafeSamplesAutomatically,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderSubmitSamplesConsentType(input)
	return &out, nil
}
