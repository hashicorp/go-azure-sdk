package synchronizationjobschema

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

type AttributeDefinitionMutability string

const (
	AttributeDefinitionMutabilityImmutable AttributeDefinitionMutability = "Immutable"
	AttributeDefinitionMutabilityReadOnly  AttributeDefinitionMutability = "ReadOnly"
	AttributeDefinitionMutabilityReadWrite AttributeDefinitionMutability = "ReadWrite"
	AttributeDefinitionMutabilityWriteOnly AttributeDefinitionMutability = "WriteOnly"
)

func PossibleValuesForAttributeDefinitionMutability() []string {
	return []string{
		string(AttributeDefinitionMutabilityImmutable),
		string(AttributeDefinitionMutabilityReadOnly),
		string(AttributeDefinitionMutabilityReadWrite),
		string(AttributeDefinitionMutabilityWriteOnly),
	}
}

func (s *AttributeDefinitionMutability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionMutability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionMutability(input string) (*AttributeDefinitionMutability, error) {
	vals := map[string]AttributeDefinitionMutability{
		"immutable": AttributeDefinitionMutabilityImmutable,
		"readonly":  AttributeDefinitionMutabilityReadOnly,
		"readwrite": AttributeDefinitionMutabilityReadWrite,
		"writeonly": AttributeDefinitionMutabilityWriteOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMutability(input)
	return &out, nil
}

type AttributeDefinitionType string

const (
	AttributeDefinitionTypeBinary    AttributeDefinitionType = "Binary"
	AttributeDefinitionTypeBoolean   AttributeDefinitionType = "Boolean"
	AttributeDefinitionTypeDateTime  AttributeDefinitionType = "DateTime"
	AttributeDefinitionTypeInteger   AttributeDefinitionType = "Integer"
	AttributeDefinitionTypeReference AttributeDefinitionType = "Reference"
	AttributeDefinitionTypeString    AttributeDefinitionType = "String"
)

func PossibleValuesForAttributeDefinitionType() []string {
	return []string{
		string(AttributeDefinitionTypeBinary),
		string(AttributeDefinitionTypeBoolean),
		string(AttributeDefinitionTypeDateTime),
		string(AttributeDefinitionTypeInteger),
		string(AttributeDefinitionTypeReference),
		string(AttributeDefinitionTypeString),
	}
}

func (s *AttributeDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionType(input string) (*AttributeDefinitionType, error) {
	vals := map[string]AttributeDefinitionType{
		"binary":    AttributeDefinitionTypeBinary,
		"boolean":   AttributeDefinitionTypeBoolean,
		"datetime":  AttributeDefinitionTypeDateTime,
		"integer":   AttributeDefinitionTypeInteger,
		"reference": AttributeDefinitionTypeReference,
		"string":    AttributeDefinitionTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionType(input)
	return &out, nil
}

type ObjectDefinitionMetadataEntryKey string

const (
	ObjectDefinitionMetadataEntryKeyBaseObjectName               ObjectDefinitionMetadataEntryKey = "BaseObjectName"
	ObjectDefinitionMetadataEntryKeyConnectorDataStorageRequired ObjectDefinitionMetadataEntryKey = "ConnectorDataStorageRequired"
	ObjectDefinitionMetadataEntryKeyExtensions                   ObjectDefinitionMetadataEntryKey = "Extensions"
	ObjectDefinitionMetadataEntryKeyIsSoftDeletionSupported      ObjectDefinitionMetadataEntryKey = "IsSoftDeletionSupported"
	ObjectDefinitionMetadataEntryKeyIsSynchronizeAllSupported    ObjectDefinitionMetadataEntryKey = "IsSynchronizeAllSupported"
	ObjectDefinitionMetadataEntryKeyPropertyNameAccountEnabled   ObjectDefinitionMetadataEntryKey = "PropertyNameAccountEnabled"
	ObjectDefinitionMetadataEntryKeyPropertyNameSoftDeleted      ObjectDefinitionMetadataEntryKey = "PropertyNameSoftDeleted"
)

func PossibleValuesForObjectDefinitionMetadataEntryKey() []string {
	return []string{
		string(ObjectDefinitionMetadataEntryKeyBaseObjectName),
		string(ObjectDefinitionMetadataEntryKeyConnectorDataStorageRequired),
		string(ObjectDefinitionMetadataEntryKeyExtensions),
		string(ObjectDefinitionMetadataEntryKeyIsSoftDeletionSupported),
		string(ObjectDefinitionMetadataEntryKeyIsSynchronizeAllSupported),
		string(ObjectDefinitionMetadataEntryKeyPropertyNameAccountEnabled),
		string(ObjectDefinitionMetadataEntryKeyPropertyNameSoftDeleted),
	}
}

func (s *ObjectDefinitionMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectDefinitionMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectDefinitionMetadataEntryKey(input string) (*ObjectDefinitionMetadataEntryKey, error) {
	vals := map[string]ObjectDefinitionMetadataEntryKey{
		"baseobjectname":               ObjectDefinitionMetadataEntryKeyBaseObjectName,
		"connectordatastoragerequired": ObjectDefinitionMetadataEntryKeyConnectorDataStorageRequired,
		"extensions":                   ObjectDefinitionMetadataEntryKeyExtensions,
		"issoftdeletionsupported":      ObjectDefinitionMetadataEntryKeyIsSoftDeletionSupported,
		"issynchronizeallsupported":    ObjectDefinitionMetadataEntryKeyIsSynchronizeAllSupported,
		"propertynameaccountenabled":   ObjectDefinitionMetadataEntryKeyPropertyNameAccountEnabled,
		"propertynamesoftdeleted":      ObjectDefinitionMetadataEntryKeyPropertyNameSoftDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectDefinitionMetadataEntryKey(input)
	return &out, nil
}
