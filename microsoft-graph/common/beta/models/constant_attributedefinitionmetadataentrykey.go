package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionMetadataEntryKey string

const (
	AttributeDefinitionMetadataEntryKeyBaseAttributeName       AttributeDefinitionMetadataEntryKey = "BaseAttributeName"
	AttributeDefinitionMetadataEntryKeyComplexObjectDefinition AttributeDefinitionMetadataEntryKey = "ComplexObjectDefinition"
	AttributeDefinitionMetadataEntryKeyIsContainer             AttributeDefinitionMetadataEntryKey = "IsContainer"
	AttributeDefinitionMetadataEntryKeyIsCustomerDefined       AttributeDefinitionMetadataEntryKey = "IsCustomerDefined"
	AttributeDefinitionMetadataEntryKeyIsDomainQualified       AttributeDefinitionMetadataEntryKey = "IsDomainQualified"
	AttributeDefinitionMetadataEntryKeyLinkPropertyNames       AttributeDefinitionMetadataEntryKey = "LinkPropertyNames"
	AttributeDefinitionMetadataEntryKeyLinkTypeName            AttributeDefinitionMetadataEntryKey = "LinkTypeName"
	AttributeDefinitionMetadataEntryKeyMaximumLength           AttributeDefinitionMetadataEntryKey = "MaximumLength"
	AttributeDefinitionMetadataEntryKeyReferencedProperty      AttributeDefinitionMetadataEntryKey = "ReferencedProperty"
)

func PossibleValuesForAttributeDefinitionMetadataEntryKey() []string {
	return []string{
		string(AttributeDefinitionMetadataEntryKeyBaseAttributeName),
		string(AttributeDefinitionMetadataEntryKeyComplexObjectDefinition),
		string(AttributeDefinitionMetadataEntryKeyIsContainer),
		string(AttributeDefinitionMetadataEntryKeyIsCustomerDefined),
		string(AttributeDefinitionMetadataEntryKeyIsDomainQualified),
		string(AttributeDefinitionMetadataEntryKeyLinkPropertyNames),
		string(AttributeDefinitionMetadataEntryKeyLinkTypeName),
		string(AttributeDefinitionMetadataEntryKeyMaximumLength),
		string(AttributeDefinitionMetadataEntryKeyReferencedProperty),
	}
}

func parseAttributeDefinitionMetadataEntryKey(input string) (*AttributeDefinitionMetadataEntryKey, error) {
	vals := map[string]AttributeDefinitionMetadataEntryKey{
		"baseattributename":       AttributeDefinitionMetadataEntryKeyBaseAttributeName,
		"complexobjectdefinition": AttributeDefinitionMetadataEntryKeyComplexObjectDefinition,
		"iscontainer":             AttributeDefinitionMetadataEntryKeyIsContainer,
		"iscustomerdefined":       AttributeDefinitionMetadataEntryKeyIsCustomerDefined,
		"isdomainqualified":       AttributeDefinitionMetadataEntryKeyIsDomainQualified,
		"linkpropertynames":       AttributeDefinitionMetadataEntryKeyLinkPropertyNames,
		"linktypename":            AttributeDefinitionMetadataEntryKeyLinkTypeName,
		"maximumlength":           AttributeDefinitionMetadataEntryKeyMaximumLength,
		"referencedproperty":      AttributeDefinitionMetadataEntryKeyReferencedProperty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMetadataEntryKey(input)
	return &out, nil
}
