package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationFindMyFiles string

const (
	Windows10GeneralConfigurationFindMyFilesdisabled      Windows10GeneralConfigurationFindMyFiles = "Disabled"
	Windows10GeneralConfigurationFindMyFilesenabled       Windows10GeneralConfigurationFindMyFiles = "Enabled"
	Windows10GeneralConfigurationFindMyFilesnotConfigured Windows10GeneralConfigurationFindMyFiles = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationFindMyFiles() []string {
	return []string{
		string(Windows10GeneralConfigurationFindMyFilesdisabled),
		string(Windows10GeneralConfigurationFindMyFilesenabled),
		string(Windows10GeneralConfigurationFindMyFilesnotConfigured),
	}
}

func parseWindows10GeneralConfigurationFindMyFiles(input string) (*Windows10GeneralConfigurationFindMyFiles, error) {
	vals := map[string]Windows10GeneralConfigurationFindMyFiles{
		"disabled":      Windows10GeneralConfigurationFindMyFilesdisabled,
		"enabled":       Windows10GeneralConfigurationFindMyFilesenabled,
		"notconfigured": Windows10GeneralConfigurationFindMyFilesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationFindMyFiles(input)
	return &out, nil
}
