package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTemplate struct {
	ApplicationId *string                         `json:"applicationId,omitempty"`
	Default       *bool                           `json:"default,omitempty"`
	Description   *string                         `json:"description,omitempty"`
	Discoverable  *bool                           `json:"discoverable,omitempty"`
	FactoryTag    *string                         `json:"factoryTag,omitempty"`
	Id            *string                         `json:"id,omitempty"`
	Metadata      *[]SynchronizationMetadataEntry `json:"metadata,omitempty"`
	ODataType     *string                         `json:"@odata.type,omitempty"`
	Schema        *SynchronizationSchema          `json:"schema,omitempty"`
}
