package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateCategorySummaryUpdateCategory string

const (
	MacOSSoftwareUpdateCategorySummaryUpdateCategory_ConfigurationDataFile MacOSSoftwareUpdateCategorySummaryUpdateCategory = "configurationDataFile"
	MacOSSoftwareUpdateCategorySummaryUpdateCategory_Critical              MacOSSoftwareUpdateCategorySummaryUpdateCategory = "critical"
	MacOSSoftwareUpdateCategorySummaryUpdateCategory_Firmware              MacOSSoftwareUpdateCategorySummaryUpdateCategory = "firmware"
	MacOSSoftwareUpdateCategorySummaryUpdateCategory_Other                 MacOSSoftwareUpdateCategorySummaryUpdateCategory = "other"
)

func PossibleValuesForMacOSSoftwareUpdateCategorySummaryUpdateCategory() []string {
	return []string{
		string(MacOSSoftwareUpdateCategorySummaryUpdateCategory_ConfigurationDataFile),
		string(MacOSSoftwareUpdateCategorySummaryUpdateCategory_Critical),
		string(MacOSSoftwareUpdateCategorySummaryUpdateCategory_Firmware),
		string(MacOSSoftwareUpdateCategorySummaryUpdateCategory_Other),
	}
}

func (s *MacOSSoftwareUpdateCategorySummaryUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateCategorySummaryUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateCategorySummaryUpdateCategory(input string) (*MacOSSoftwareUpdateCategorySummaryUpdateCategory, error) {
	vals := map[string]MacOSSoftwareUpdateCategorySummaryUpdateCategory{
		"configurationdatafile": MacOSSoftwareUpdateCategorySummaryUpdateCategory_ConfigurationDataFile,
		"critical":              MacOSSoftwareUpdateCategorySummaryUpdateCategory_Critical,
		"firmware":              MacOSSoftwareUpdateCategorySummaryUpdateCategory_Firmware,
		"other":                 MacOSSoftwareUpdateCategorySummaryUpdateCategory_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateCategorySummaryUpdateCategory(input)
	return &out, nil
}
