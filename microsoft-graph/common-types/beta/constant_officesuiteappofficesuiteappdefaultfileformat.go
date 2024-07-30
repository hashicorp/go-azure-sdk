package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppOfficeSuiteAppDefaultFileFormat string

const (
	OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_NotConfigured            OfficeSuiteAppOfficeSuiteAppDefaultFileFormat = "notConfigured"
	OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_OfficeOpenDocumentFormat OfficeSuiteAppOfficeSuiteAppDefaultFileFormat = "officeOpenDocumentFormat"
	OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_OfficeOpenXMLFormat      OfficeSuiteAppOfficeSuiteAppDefaultFileFormat = "officeOpenXMLFormat"
)

func PossibleValuesForOfficeSuiteAppOfficeSuiteAppDefaultFileFormat() []string {
	return []string{
		string(OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_NotConfigured),
		string(OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_OfficeOpenDocumentFormat),
		string(OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_OfficeOpenXMLFormat),
	}
}

func (s *OfficeSuiteAppOfficeSuiteAppDefaultFileFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteAppOfficeSuiteAppDefaultFileFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteAppOfficeSuiteAppDefaultFileFormat(input string) (*OfficeSuiteAppOfficeSuiteAppDefaultFileFormat, error) {
	vals := map[string]OfficeSuiteAppOfficeSuiteAppDefaultFileFormat{
		"notconfigured":            OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_NotConfigured,
		"officeopendocumentformat": OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_OfficeOpenDocumentFormat,
		"officeopenxmlformat":      OfficeSuiteAppOfficeSuiteAppDefaultFileFormat_OfficeOpenXMLFormat,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppOfficeSuiteAppDefaultFileFormat(input)
	return &out, nil
}
