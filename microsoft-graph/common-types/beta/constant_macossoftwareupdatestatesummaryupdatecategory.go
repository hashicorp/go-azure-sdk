package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateStateSummaryUpdateCategory string

const (
	MacOSSoftwareUpdateStateSummaryUpdateCategory_ConfigurationDataFile MacOSSoftwareUpdateStateSummaryUpdateCategory = "configurationDataFile"
	MacOSSoftwareUpdateStateSummaryUpdateCategory_Critical              MacOSSoftwareUpdateStateSummaryUpdateCategory = "critical"
	MacOSSoftwareUpdateStateSummaryUpdateCategory_Firmware              MacOSSoftwareUpdateStateSummaryUpdateCategory = "firmware"
	MacOSSoftwareUpdateStateSummaryUpdateCategory_Other                 MacOSSoftwareUpdateStateSummaryUpdateCategory = "other"
)

func PossibleValuesForMacOSSoftwareUpdateStateSummaryUpdateCategory() []string {
	return []string{
		string(MacOSSoftwareUpdateStateSummaryUpdateCategory_ConfigurationDataFile),
		string(MacOSSoftwareUpdateStateSummaryUpdateCategory_Critical),
		string(MacOSSoftwareUpdateStateSummaryUpdateCategory_Firmware),
		string(MacOSSoftwareUpdateStateSummaryUpdateCategory_Other),
	}
}

func (s *MacOSSoftwareUpdateStateSummaryUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateStateSummaryUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateStateSummaryUpdateCategory(input string) (*MacOSSoftwareUpdateStateSummaryUpdateCategory, error) {
	vals := map[string]MacOSSoftwareUpdateStateSummaryUpdateCategory{
		"configurationdatafile": MacOSSoftwareUpdateStateSummaryUpdateCategory_ConfigurationDataFile,
		"critical":              MacOSSoftwareUpdateStateSummaryUpdateCategory_Critical,
		"firmware":              MacOSSoftwareUpdateStateSummaryUpdateCategory_Firmware,
		"other":                 MacOSSoftwareUpdateStateSummaryUpdateCategory_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateStateSummaryUpdateCategory(input)
	return &out, nil
}
