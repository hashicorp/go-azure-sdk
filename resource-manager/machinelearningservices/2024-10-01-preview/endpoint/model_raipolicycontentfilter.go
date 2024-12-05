package endpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiPolicyContentFilter struct {
	AllowedContentLevel *AllowedContentLevel    `json:"allowedContentLevel,omitempty"`
	Blocking            *bool                   `json:"blocking,omitempty"`
	Enabled             *bool                   `json:"enabled,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	Source              *RaiPolicyContentSource `json:"source,omitempty"`
}
