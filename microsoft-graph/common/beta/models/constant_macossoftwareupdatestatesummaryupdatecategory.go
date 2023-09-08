package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateStateSummaryUpdateCategory string

const (
	MacOSSoftwareUpdateStateSummaryUpdateCategoryconfigurationDataFile MacOSSoftwareUpdateStateSummaryUpdateCategory = "ConfigurationDataFile"
	MacOSSoftwareUpdateStateSummaryUpdateCategorycritical              MacOSSoftwareUpdateStateSummaryUpdateCategory = "Critical"
	MacOSSoftwareUpdateStateSummaryUpdateCategoryfirmware              MacOSSoftwareUpdateStateSummaryUpdateCategory = "Firmware"
	MacOSSoftwareUpdateStateSummaryUpdateCategoryother                 MacOSSoftwareUpdateStateSummaryUpdateCategory = "Other"
)

func PossibleValuesForMacOSSoftwareUpdateStateSummaryUpdateCategory() []string {
	return []string{
		string(MacOSSoftwareUpdateStateSummaryUpdateCategoryconfigurationDataFile),
		string(MacOSSoftwareUpdateStateSummaryUpdateCategorycritical),
		string(MacOSSoftwareUpdateStateSummaryUpdateCategoryfirmware),
		string(MacOSSoftwareUpdateStateSummaryUpdateCategoryother),
	}
}

func parseMacOSSoftwareUpdateStateSummaryUpdateCategory(input string) (*MacOSSoftwareUpdateStateSummaryUpdateCategory, error) {
	vals := map[string]MacOSSoftwareUpdateStateSummaryUpdateCategory{
		"configurationdatafile": MacOSSoftwareUpdateStateSummaryUpdateCategoryconfigurationDataFile,
		"critical":              MacOSSoftwareUpdateStateSummaryUpdateCategorycritical,
		"firmware":              MacOSSoftwareUpdateStateSummaryUpdateCategoryfirmware,
		"other":                 MacOSSoftwareUpdateStateSummaryUpdateCategoryother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateStateSummaryUpdateCategory(input)
	return &out, nil
}
