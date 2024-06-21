package catalogs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CatalogProperties struct {
	AdoGit            *GitCatalog        `json:"adoGit,omitempty"`
	GitHub            *GitCatalog        `json:"gitHub,omitempty"`
	LastSyncTime      *string            `json:"lastSyncTime,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	SyncState         *CatalogSyncState  `json:"syncState,omitempty"`
}
