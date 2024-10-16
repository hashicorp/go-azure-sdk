package policymetadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyMetadataSlimProperties struct {
	AdditionalContentURL *string      `json:"additionalContentUrl,omitempty"`
	Category             *string      `json:"category,omitempty"`
	Metadata             *interface{} `json:"metadata,omitempty"`
	MetadataId           *string      `json:"metadataId,omitempty"`
	Owner                *string      `json:"owner,omitempty"`
	Title                *string      `json:"title,omitempty"`
}
