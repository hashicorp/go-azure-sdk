package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationCellularData string

const (
	Windows10GeneralConfigurationCellularDataallowed       Windows10GeneralConfigurationCellularData = "Allowed"
	Windows10GeneralConfigurationCellularDatablocked       Windows10GeneralConfigurationCellularData = "Blocked"
	Windows10GeneralConfigurationCellularDatanotConfigured Windows10GeneralConfigurationCellularData = "NotConfigured"
	Windows10GeneralConfigurationCellularDatarequired      Windows10GeneralConfigurationCellularData = "Required"
)

func PossibleValuesForWindows10GeneralConfigurationCellularData() []string {
	return []string{
		string(Windows10GeneralConfigurationCellularDataallowed),
		string(Windows10GeneralConfigurationCellularDatablocked),
		string(Windows10GeneralConfigurationCellularDatanotConfigured),
		string(Windows10GeneralConfigurationCellularDatarequired),
	}
}

func parseWindows10GeneralConfigurationCellularData(input string) (*Windows10GeneralConfigurationCellularData, error) {
	vals := map[string]Windows10GeneralConfigurationCellularData{
		"allowed":       Windows10GeneralConfigurationCellularDataallowed,
		"blocked":       Windows10GeneralConfigurationCellularDatablocked,
		"notconfigured": Windows10GeneralConfigurationCellularDatanotConfigured,
		"required":      Windows10GeneralConfigurationCellularDatarequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationCellularData(input)
	return &out, nil
}
