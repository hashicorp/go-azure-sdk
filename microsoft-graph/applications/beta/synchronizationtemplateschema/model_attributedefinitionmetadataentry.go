package synchronizationtemplateschema

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionMetadataEntry struct {
	Key       *AttributeDefinitionMetadataEntryKey `json:"key,omitempty"`
	ODataType *string                              `json:"@odata.type,omitempty"`
	Value     *string                              `json:"value,omitempty"`
}
