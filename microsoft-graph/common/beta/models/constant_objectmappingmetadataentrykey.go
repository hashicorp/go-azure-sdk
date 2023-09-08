package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingMetadataEntryKey string

const (
	ObjectMappingMetadataEntryKeyDisableMonitoringForChanges ObjectMappingMetadataEntryKey = "DisableMonitoringForChanges"
	ObjectMappingMetadataEntryKeyDisposition                 ObjectMappingMetadataEntryKey = "Disposition"
	ObjectMappingMetadataEntryKeyEscrowBehavior              ObjectMappingMetadataEntryKey = "EscrowBehavior"
	ObjectMappingMetadataEntryKeyExcludeFromReporting        ObjectMappingMetadataEntryKey = "ExcludeFromReporting"
	ObjectMappingMetadataEntryKeyIsCustomerDefined           ObjectMappingMetadataEntryKey = "IsCustomerDefined"
	ObjectMappingMetadataEntryKeyOriginalJoiningProperty     ObjectMappingMetadataEntryKey = "OriginalJoiningProperty"
	ObjectMappingMetadataEntryKeyUnsynchronized              ObjectMappingMetadataEntryKey = "Unsynchronized"
)

func PossibleValuesForObjectMappingMetadataEntryKey() []string {
	return []string{
		string(ObjectMappingMetadataEntryKeyDisableMonitoringForChanges),
		string(ObjectMappingMetadataEntryKeyDisposition),
		string(ObjectMappingMetadataEntryKeyEscrowBehavior),
		string(ObjectMappingMetadataEntryKeyExcludeFromReporting),
		string(ObjectMappingMetadataEntryKeyIsCustomerDefined),
		string(ObjectMappingMetadataEntryKeyOriginalJoiningProperty),
		string(ObjectMappingMetadataEntryKeyUnsynchronized),
	}
}

func parseObjectMappingMetadataEntryKey(input string) (*ObjectMappingMetadataEntryKey, error) {
	vals := map[string]ObjectMappingMetadataEntryKey{
		"disablemonitoringforchanges": ObjectMappingMetadataEntryKeyDisableMonitoringForChanges,
		"disposition":                 ObjectMappingMetadataEntryKeyDisposition,
		"escrowbehavior":              ObjectMappingMetadataEntryKeyEscrowBehavior,
		"excludefromreporting":        ObjectMappingMetadataEntryKeyExcludeFromReporting,
		"iscustomerdefined":           ObjectMappingMetadataEntryKeyIsCustomerDefined,
		"originaljoiningproperty":     ObjectMappingMetadataEntryKeyOriginalJoiningProperty,
		"unsynchronized":              ObjectMappingMetadataEntryKeyUnsynchronized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingMetadataEntryKey(input)
	return &out, nil
}
