package labs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabAnnouncementProperties struct {
	Enabled           *EnableStatus `json:"enabled,omitempty"`
	ExpirationDate    *string       `json:"expirationDate,omitempty"`
	Expired           *bool         `json:"expired,omitempty"`
	Markdown          *string       `json:"markdown,omitempty"`
	ProvisioningState *string       `json:"provisioningState,omitempty"`
	Title             *string       `json:"title,omitempty"`
	UniqueIdentifier  *string       `json:"uniqueIdentifier,omitempty"`
}
