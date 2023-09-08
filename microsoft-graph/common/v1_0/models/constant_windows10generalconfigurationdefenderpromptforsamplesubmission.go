package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderPromptForSampleSubmission string

const (
	Windows10GeneralConfigurationDefenderPromptForSampleSubmissionalwaysPrompt                    Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "AlwaysPrompt"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmissionneverSendData                   Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "NeverSendData"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmissionpromptBeforeSendingPersonalData Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "PromptBeforeSendingPersonalData"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmissionsendAllDataWithoutPrompting     Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "SendAllDataWithoutPrompting"
	Windows10GeneralConfigurationDefenderPromptForSampleSubmissionuserDefined                     Windows10GeneralConfigurationDefenderPromptForSampleSubmission = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderPromptForSampleSubmission() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmissionalwaysPrompt),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmissionneverSendData),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmissionpromptBeforeSendingPersonalData),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmissionsendAllDataWithoutPrompting),
		string(Windows10GeneralConfigurationDefenderPromptForSampleSubmissionuserDefined),
	}
}

func parseWindows10GeneralConfigurationDefenderPromptForSampleSubmission(input string) (*Windows10GeneralConfigurationDefenderPromptForSampleSubmission, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderPromptForSampleSubmission{
		"alwaysprompt":                    Windows10GeneralConfigurationDefenderPromptForSampleSubmissionalwaysPrompt,
		"neversenddata":                   Windows10GeneralConfigurationDefenderPromptForSampleSubmissionneverSendData,
		"promptbeforesendingpersonaldata": Windows10GeneralConfigurationDefenderPromptForSampleSubmissionpromptBeforeSendingPersonalData,
		"sendalldatawithoutprompting":     Windows10GeneralConfigurationDefenderPromptForSampleSubmissionsendAllDataWithoutPrompting,
		"userdefined":                     Windows10GeneralConfigurationDefenderPromptForSampleSubmissionuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderPromptForSampleSubmission(input)
	return &out, nil
}
