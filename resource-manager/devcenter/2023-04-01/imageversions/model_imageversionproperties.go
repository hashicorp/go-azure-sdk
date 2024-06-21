package imageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageVersionProperties struct {
	ExcludeFromLatest   *bool              `json:"excludeFromLatest,omitempty"`
	Name                *string            `json:"name,omitempty"`
	OsDiskImageSizeInGb *int64             `json:"osDiskImageSizeInGb,omitempty"`
	ProvisioningState   *ProvisioningState `json:"provisioningState,omitempty"`
	PublishedDate       *string            `json:"publishedDate,omitempty"`
}
