package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingMetadataEntryKey string

const (
	ObjectMappingMetadataEntryKey_DisableMonitoringForChanges ObjectMappingMetadataEntryKey = "DisableMonitoringForChanges"
	ObjectMappingMetadataEntryKey_Disposition                 ObjectMappingMetadataEntryKey = "Disposition"
	ObjectMappingMetadataEntryKey_EscrowBehavior              ObjectMappingMetadataEntryKey = "EscrowBehavior"
	ObjectMappingMetadataEntryKey_ExcludeFromReporting        ObjectMappingMetadataEntryKey = "ExcludeFromReporting"
	ObjectMappingMetadataEntryKey_IsCustomerDefined           ObjectMappingMetadataEntryKey = "IsCustomerDefined"
	ObjectMappingMetadataEntryKey_OriginalJoiningProperty     ObjectMappingMetadataEntryKey = "OriginalJoiningProperty"
	ObjectMappingMetadataEntryKey_Unsynchronized              ObjectMappingMetadataEntryKey = "Unsynchronized"
)

func PossibleValuesForObjectMappingMetadataEntryKey() []string {
	return []string{
		string(ObjectMappingMetadataEntryKey_DisableMonitoringForChanges),
		string(ObjectMappingMetadataEntryKey_Disposition),
		string(ObjectMappingMetadataEntryKey_EscrowBehavior),
		string(ObjectMappingMetadataEntryKey_ExcludeFromReporting),
		string(ObjectMappingMetadataEntryKey_IsCustomerDefined),
		string(ObjectMappingMetadataEntryKey_OriginalJoiningProperty),
		string(ObjectMappingMetadataEntryKey_Unsynchronized),
	}
}

func (s *ObjectMappingMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectMappingMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectMappingMetadataEntryKey(input string) (*ObjectMappingMetadataEntryKey, error) {
	vals := map[string]ObjectMappingMetadataEntryKey{
		"disablemonitoringforchanges": ObjectMappingMetadataEntryKey_DisableMonitoringForChanges,
		"disposition":                 ObjectMappingMetadataEntryKey_Disposition,
		"escrowbehavior":              ObjectMappingMetadataEntryKey_EscrowBehavior,
		"excludefromreporting":        ObjectMappingMetadataEntryKey_ExcludeFromReporting,
		"iscustomerdefined":           ObjectMappingMetadataEntryKey_IsCustomerDefined,
		"originaljoiningproperty":     ObjectMappingMetadataEntryKey_OriginalJoiningProperty,
		"unsynchronized":              ObjectMappingMetadataEntryKey_Unsynchronized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingMetadataEntryKey(input)
	return &out, nil
}
