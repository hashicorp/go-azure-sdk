package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentContentProperties struct {
	CurrentLabel             *CurrentLabel              `json:"currentLabel,omitempty"`
	DiscoveredSensitiveTypes *[]DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`
	Extensions               *[]string                  `json:"extensions,omitempty"`
	LastModifiedBy           *string                    `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime     *string                    `json:"lastModifiedDateTime,omitempty"`
	Metadata                 *ContentMetadata           `json:"metadata,omitempty"`
	ODataType                *string                    `json:"@odata.type,omitempty"`
}
