package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity string

const (
	WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeveritycritical  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity = "Critical"
	WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverityimportant WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity = "Important"
	WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeveritymoderate  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity = "Moderate"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeveritycritical),
		string(WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverityimportant),
		string(WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeveritymoderate),
	}
}

func parseWindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity(input string) (*WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity{
		"critical":  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeveritycritical,
		"important": WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverityimportant,
		"moderate":  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeveritymoderate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity(input)
	return &out, nil
}
