package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDefinitionMetadataEntry struct {
	Key       *ObjectDefinitionMetadataEntryKey `json:"key,omitempty"`
	ODataType *string                           `json:"@odata.type,omitempty"`
	Value     *string                           `json:"value,omitempty"`
}
