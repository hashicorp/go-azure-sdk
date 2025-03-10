package v2workspaceconnectionresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiBlocklistItemBulkRequest struct {
	Name       *string                     `json:"name,omitempty"`
	Properties *RaiBlocklistItemProperties `json:"properties,omitempty"`
}
