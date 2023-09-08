package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLearningSummaryApplicationType string

const (
	WindowsInformationProtectionAppLearningSummaryApplicationTypedesktop   WindowsInformationProtectionAppLearningSummaryApplicationType = "Desktop"
	WindowsInformationProtectionAppLearningSummaryApplicationTypeuniversal WindowsInformationProtectionAppLearningSummaryApplicationType = "Universal"
)

func PossibleValuesForWindowsInformationProtectionAppLearningSummaryApplicationType() []string {
	return []string{
		string(WindowsInformationProtectionAppLearningSummaryApplicationTypedesktop),
		string(WindowsInformationProtectionAppLearningSummaryApplicationTypeuniversal),
	}
}

func parseWindowsInformationProtectionAppLearningSummaryApplicationType(input string) (*WindowsInformationProtectionAppLearningSummaryApplicationType, error) {
	vals := map[string]WindowsInformationProtectionAppLearningSummaryApplicationType{
		"desktop":   WindowsInformationProtectionAppLearningSummaryApplicationTypedesktop,
		"universal": WindowsInformationProtectionAppLearningSummaryApplicationTypeuniversal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionAppLearningSummaryApplicationType(input)
	return &out, nil
}
