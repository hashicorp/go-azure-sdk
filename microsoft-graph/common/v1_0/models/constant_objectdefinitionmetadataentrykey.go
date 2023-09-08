package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
