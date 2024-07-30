package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDefinitionMetadataEntryKey string

const (
	ObjectDefinitionMetadataEntryKey_BaseObjectName               ObjectDefinitionMetadataEntryKey = "BaseObjectName"
	ObjectDefinitionMetadataEntryKey_ConnectorDataStorageRequired ObjectDefinitionMetadataEntryKey = "ConnectorDataStorageRequired"
	ObjectDefinitionMetadataEntryKey_Extensions                   ObjectDefinitionMetadataEntryKey = "Extensions"
	ObjectDefinitionMetadataEntryKey_IsSoftDeletionSupported      ObjectDefinitionMetadataEntryKey = "IsSoftDeletionSupported"
	ObjectDefinitionMetadataEntryKey_IsSynchronizeAllSupported    ObjectDefinitionMetadataEntryKey = "IsSynchronizeAllSupported"
	ObjectDefinitionMetadataEntryKey_PropertyNameAccountEnabled   ObjectDefinitionMetadataEntryKey = "PropertyNameAccountEnabled"
	ObjectDefinitionMetadataEntryKey_PropertyNameSoftDeleted      ObjectDefinitionMetadataEntryKey = "PropertyNameSoftDeleted"
)

func PossibleValuesForObjectDefinitionMetadataEntryKey() []string {
	return []string{
		string(ObjectDefinitionMetadataEntryKey_BaseObjectName),
		string(ObjectDefinitionMetadataEntryKey_ConnectorDataStorageRequired),
		string(ObjectDefinitionMetadataEntryKey_Extensions),
		string(ObjectDefinitionMetadataEntryKey_IsSoftDeletionSupported),
		string(ObjectDefinitionMetadataEntryKey_IsSynchronizeAllSupported),
		string(ObjectDefinitionMetadataEntryKey_PropertyNameAccountEnabled),
		string(ObjectDefinitionMetadataEntryKey_PropertyNameSoftDeleted),
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
		"baseobjectname":               ObjectDefinitionMetadataEntryKey_BaseObjectName,
		"connectordatastoragerequired": ObjectDefinitionMetadataEntryKey_ConnectorDataStorageRequired,
		"extensions":                   ObjectDefinitionMetadataEntryKey_Extensions,
		"issoftdeletionsupported":      ObjectDefinitionMetadataEntryKey_IsSoftDeletionSupported,
		"issynchronizeallsupported":    ObjectDefinitionMetadataEntryKey_IsSynchronizeAllSupported,
		"propertynameaccountenabled":   ObjectDefinitionMetadataEntryKey_PropertyNameAccountEnabled,
		"propertynamesoftdeleted":      ObjectDefinitionMetadataEntryKey_PropertyNameSoftDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectDefinitionMetadataEntryKey(input)
	return &out, nil
}
