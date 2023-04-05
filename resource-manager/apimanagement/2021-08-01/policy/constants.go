package policy

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
