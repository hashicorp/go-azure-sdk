package policy

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyContentFormat string

const (
	PolicyContentFormatRawxml             PolicyContentFormat = "rawxml"
	PolicyContentFormatRawxmlNegativelink PolicyContentFormat = "rawxml-link"
	PolicyContentFormatXml                PolicyContentFormat = "xml"
	PolicyContentFormatXmlNegativelink    PolicyContentFormat = "xml-link"
)

func PossibleValuesForPolicyContentFormat() []string {
	return []string{
		string(PolicyContentFormatRawxml),
		string(PolicyContentFormatRawxmlNegativelink),
		string(PolicyContentFormatXml),
		string(PolicyContentFormatXmlNegativelink),
	}
}

func parsePolicyContentFormat(input string) (*PolicyContentFormat, error) {
	vals := map[string]PolicyContentFormat{
		"rawxml":      PolicyContentFormatRawxml,
		"rawxml-link": PolicyContentFormatRawxmlNegativelink,
		"xml":         PolicyContentFormatXml,
		"xml-link":    PolicyContentFormatXmlNegativelink,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyContentFormat(input)
	return &out, nil
}

type PolicyExportFormat string

const (
	PolicyExportFormatRawxml PolicyExportFormat = "rawxml"
	PolicyExportFormatXml    PolicyExportFormat = "xml"
)

func PossibleValuesForPolicyExportFormat() []string {
	return []string{
		string(PolicyExportFormatRawxml),
		string(PolicyExportFormatXml),
	}
}

func parsePolicyExportFormat(input string) (*PolicyExportFormat, error) {
	vals := map[string]PolicyExportFormat{
		"rawxml": PolicyExportFormatRawxml,
		"xml":    PolicyExportFormatXml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyExportFormat(input)
	return &out, nil
}
