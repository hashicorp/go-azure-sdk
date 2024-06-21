package watchlistitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatchlistItemProperties struct {
	Created           *string      `json:"created,omitempty"`
	CreatedBy         *UserInfo    `json:"createdBy,omitempty"`
	EntityMapping     *interface{} `json:"entityMapping,omitempty"`
	IsDeleted         *bool        `json:"isDeleted,omitempty"`
	ItemsKeyValue     interface{}  `json:"itemsKeyValue"`
	TenantId          *string      `json:"tenantId,omitempty"`
	Updated           *string      `json:"updated,omitempty"`
	UpdatedBy         *UserInfo    `json:"updatedBy,omitempty"`
	WatchlistItemId   *string      `json:"watchlistItemId,omitempty"`
	WatchlistItemType *string      `json:"watchlistItemType,omitempty"`
}
