package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportSuspiciousActivitySettingsState string

const (
	ReportSuspiciousActivitySettingsStatedefault  ReportSuspiciousActivitySettingsState = "Default"
	ReportSuspiciousActivitySettingsStatedisabled ReportSuspiciousActivitySettingsState = "Disabled"
	ReportSuspiciousActivitySettingsStateenabled  ReportSuspiciousActivitySettingsState = "Enabled"
)

func PossibleValuesForReportSuspiciousActivitySettingsState() []string {
	return []string{
		string(ReportSuspiciousActivitySettingsStatedefault),
		string(ReportSuspiciousActivitySettingsStatedisabled),
		string(ReportSuspiciousActivitySettingsStateenabled),
	}
}

func parseReportSuspiciousActivitySettingsState(input string) (*ReportSuspiciousActivitySettingsState, error) {
	vals := map[string]ReportSuspiciousActivitySettingsState{
		"default":  ReportSuspiciousActivitySettingsStatedefault,
		"disabled": ReportSuspiciousActivitySettingsStatedisabled,
		"enabled":  ReportSuspiciousActivitySettingsStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReportSuspiciousActivitySettingsState(input)
	return &out, nil
}
