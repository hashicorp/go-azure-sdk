package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookResourceInfo struct {
	Fqdn                     *string                   `json:"fqdn,omitempty"`
	IsPrivateLinkEnabled     *bool                     `json:"isPrivateLinkEnabled,omitempty"`
	NotebookPreparationError *NotebookPreparationError `json:"notebookPreparationError,omitempty"`
	ResourceId               *string                   `json:"resourceId,omitempty"`
}
