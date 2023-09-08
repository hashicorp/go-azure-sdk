package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryDefinitionDiscoverabilities string

const (
	DirectoryDefinitionDiscoverabilitiesAttributeDataTypes  DirectoryDefinitionDiscoverabilities = "AttributeDataTypes"
	DirectoryDefinitionDiscoverabilitiesAttributeNames      DirectoryDefinitionDiscoverabilities = "AttributeNames"
	DirectoryDefinitionDiscoverabilitiesAttributeReadOnly   DirectoryDefinitionDiscoverabilities = "AttributeReadOnly"
	DirectoryDefinitionDiscoverabilitiesNone                DirectoryDefinitionDiscoverabilities = "None"
	DirectoryDefinitionDiscoverabilitiesReferenceAttributes DirectoryDefinitionDiscoverabilities = "ReferenceAttributes"
)

func PossibleValuesForDirectoryDefinitionDiscoverabilities() []string {
	return []string{
		string(DirectoryDefinitionDiscoverabilitiesAttributeDataTypes),
		string(DirectoryDefinitionDiscoverabilitiesAttributeNames),
		string(DirectoryDefinitionDiscoverabilitiesAttributeReadOnly),
		string(DirectoryDefinitionDiscoverabilitiesNone),
		string(DirectoryDefinitionDiscoverabilitiesReferenceAttributes),
	}
}

func parseDirectoryDefinitionDiscoverabilities(input string) (*DirectoryDefinitionDiscoverabilities, error) {
	vals := map[string]DirectoryDefinitionDiscoverabilities{
		"attributedatatypes":  DirectoryDefinitionDiscoverabilitiesAttributeDataTypes,
		"attributenames":      DirectoryDefinitionDiscoverabilitiesAttributeNames,
		"attributereadonly":   DirectoryDefinitionDiscoverabilitiesAttributeReadOnly,
		"none":                DirectoryDefinitionDiscoverabilitiesNone,
		"referenceattributes": DirectoryDefinitionDiscoverabilitiesReferenceAttributes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DirectoryDefinitionDiscoverabilities(input)
	return &out, nil
}
