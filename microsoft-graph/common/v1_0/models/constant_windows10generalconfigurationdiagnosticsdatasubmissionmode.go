package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDiagnosticsDataSubmissionMode string

const (
	Windows10GeneralConfigurationDiagnosticsDataSubmissionModebasic       Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "Basic"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionModeenhanced    Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "Enhanced"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionModefull        Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "Full"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionModenone        Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "None"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionModeuserDefined Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDiagnosticsDataSubmissionMode() []string {
	return []string{
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionModebasic),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionModeenhanced),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionModefull),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionModenone),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionModeuserDefined),
	}
}

func parseWindows10GeneralConfigurationDiagnosticsDataSubmissionMode(input string) (*Windows10GeneralConfigurationDiagnosticsDataSubmissionMode, error) {
	vals := map[string]Windows10GeneralConfigurationDiagnosticsDataSubmissionMode{
		"basic":       Windows10GeneralConfigurationDiagnosticsDataSubmissionModebasic,
		"enhanced":    Windows10GeneralConfigurationDiagnosticsDataSubmissionModeenhanced,
		"full":        Windows10GeneralConfigurationDiagnosticsDataSubmissionModefull,
		"none":        Windows10GeneralConfigurationDiagnosticsDataSubmissionModenone,
		"userdefined": Windows10GeneralConfigurationDiagnosticsDataSubmissionModeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDiagnosticsDataSubmissionMode(input)
	return &out, nil
}
