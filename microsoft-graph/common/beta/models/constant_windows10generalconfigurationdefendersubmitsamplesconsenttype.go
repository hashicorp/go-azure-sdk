package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderSubmitSamplesConsentType string

const (
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypealwaysPrompt                 Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "AlwaysPrompt"
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypeneverSend                    Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "NeverSend"
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypesendAllSamplesAutomatically  Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "SendAllSamplesAutomatically"
	Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypesendSafeSamplesAutomatically Windows10GeneralConfigurationDefenderSubmitSamplesConsentType = "SendSafeSamplesAutomatically"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderSubmitSamplesConsentType() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypealwaysPrompt),
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypeneverSend),
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypesendAllSamplesAutomatically),
		string(Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypesendSafeSamplesAutomatically),
	}
}

func parseWindows10GeneralConfigurationDefenderSubmitSamplesConsentType(input string) (*Windows10GeneralConfigurationDefenderSubmitSamplesConsentType, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderSubmitSamplesConsentType{
		"alwaysprompt":                 Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypealwaysPrompt,
		"neversend":                    Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypeneverSend,
		"sendallsamplesautomatically":  Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypesendAllSamplesAutomatically,
		"sendsafesamplesautomatically": Windows10GeneralConfigurationDefenderSubmitSamplesConsentTypesendSafeSamplesAutomatically,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderSubmitSamplesConsentType(input)
	return &out, nil
}
