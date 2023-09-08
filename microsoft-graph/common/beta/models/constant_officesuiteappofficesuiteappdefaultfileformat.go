package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppOfficeSuiteAppDefaultFileFormat string

const (
	OfficeSuiteAppOfficeSuiteAppDefaultFileFormatnotConfigured            OfficeSuiteAppOfficeSuiteAppDefaultFileFormat = "NotConfigured"
	OfficeSuiteAppOfficeSuiteAppDefaultFileFormatofficeOpenDocumentFormat OfficeSuiteAppOfficeSuiteAppDefaultFileFormat = "OfficeOpenDocumentFormat"
	OfficeSuiteAppOfficeSuiteAppDefaultFileFormatofficeOpenXMLFormat      OfficeSuiteAppOfficeSuiteAppDefaultFileFormat = "OfficeOpenXMLFormat"
)

func PossibleValuesForOfficeSuiteAppOfficeSuiteAppDefaultFileFormat() []string {
	return []string{
		string(OfficeSuiteAppOfficeSuiteAppDefaultFileFormatnotConfigured),
		string(OfficeSuiteAppOfficeSuiteAppDefaultFileFormatofficeOpenDocumentFormat),
		string(OfficeSuiteAppOfficeSuiteAppDefaultFileFormatofficeOpenXMLFormat),
	}
}

func parseOfficeSuiteAppOfficeSuiteAppDefaultFileFormat(input string) (*OfficeSuiteAppOfficeSuiteAppDefaultFileFormat, error) {
	vals := map[string]OfficeSuiteAppOfficeSuiteAppDefaultFileFormat{
		"notconfigured":            OfficeSuiteAppOfficeSuiteAppDefaultFileFormatnotConfigured,
		"officeopendocumentformat": OfficeSuiteAppOfficeSuiteAppDefaultFileFormatofficeOpenDocumentFormat,
		"officeopenxmlformat":      OfficeSuiteAppOfficeSuiteAppDefaultFileFormatofficeOpenXMLFormat,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppOfficeSuiteAppDefaultFileFormat(input)
	return &out, nil
}
