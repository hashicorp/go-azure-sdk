package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionMetadataEntryKey string

const (
	AttributeDefinitionMetadataEntryKey_BaseAttributeName       AttributeDefinitionMetadataEntryKey = "BaseAttributeName"
	AttributeDefinitionMetadataEntryKey_ComplexObjectDefinition AttributeDefinitionMetadataEntryKey = "ComplexObjectDefinition"
	AttributeDefinitionMetadataEntryKey_IsContainer             AttributeDefinitionMetadataEntryKey = "IsContainer"
	AttributeDefinitionMetadataEntryKey_IsCustomerDefined       AttributeDefinitionMetadataEntryKey = "IsCustomerDefined"
	AttributeDefinitionMetadataEntryKey_IsDomainQualified       AttributeDefinitionMetadataEntryKey = "IsDomainQualified"
	AttributeDefinitionMetadataEntryKey_LinkPropertyNames       AttributeDefinitionMetadataEntryKey = "LinkPropertyNames"
	AttributeDefinitionMetadataEntryKey_LinkTypeName            AttributeDefinitionMetadataEntryKey = "LinkTypeName"
	AttributeDefinitionMetadataEntryKey_MaximumLength           AttributeDefinitionMetadataEntryKey = "MaximumLength"
	AttributeDefinitionMetadataEntryKey_ReferencedProperty      AttributeDefinitionMetadataEntryKey = "ReferencedProperty"
)

func PossibleValuesForAttributeDefinitionMetadataEntryKey() []string {
	return []string{
		string(AttributeDefinitionMetadataEntryKey_BaseAttributeName),
		string(AttributeDefinitionMetadataEntryKey_ComplexObjectDefinition),
		string(AttributeDefinitionMetadataEntryKey_IsContainer),
		string(AttributeDefinitionMetadataEntryKey_IsCustomerDefined),
		string(AttributeDefinitionMetadataEntryKey_IsDomainQualified),
		string(AttributeDefinitionMetadataEntryKey_LinkPropertyNames),
		string(AttributeDefinitionMetadataEntryKey_LinkTypeName),
		string(AttributeDefinitionMetadataEntryKey_MaximumLength),
		string(AttributeDefinitionMetadataEntryKey_ReferencedProperty),
	}
}

func (s *AttributeDefinitionMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionMetadataEntryKey(input string) (*AttributeDefinitionMetadataEntryKey, error) {
	vals := map[string]AttributeDefinitionMetadataEntryKey{
		"baseattributename":       AttributeDefinitionMetadataEntryKey_BaseAttributeName,
		"complexobjectdefinition": AttributeDefinitionMetadataEntryKey_ComplexObjectDefinition,
		"iscontainer":             AttributeDefinitionMetadataEntryKey_IsContainer,
		"iscustomerdefined":       AttributeDefinitionMetadataEntryKey_IsCustomerDefined,
		"isdomainqualified":       AttributeDefinitionMetadataEntryKey_IsDomainQualified,
		"linkpropertynames":       AttributeDefinitionMetadataEntryKey_LinkPropertyNames,
		"linktypename":            AttributeDefinitionMetadataEntryKey_LinkTypeName,
		"maximumlength":           AttributeDefinitionMetadataEntryKey_MaximumLength,
		"referencedproperty":      AttributeDefinitionMetadataEntryKey_ReferencedProperty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMetadataEntryKey(input)
	return &out, nil
}
