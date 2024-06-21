package watchlists

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatchlistProperties struct {
	ContentType         *string     `json:"contentType,omitempty"`
	Created             *string     `json:"created,omitempty"`
	CreatedBy           *UserInfo   `json:"createdBy,omitempty"`
	DefaultDuration     *string     `json:"defaultDuration,omitempty"`
	Description         *string     `json:"description,omitempty"`
	DisplayName         string      `json:"displayName"`
	IsDeleted           *bool       `json:"isDeleted,omitempty"`
	ItemsSearchKey      string      `json:"itemsSearchKey"`
	Labels              *[]string   `json:"labels,omitempty"`
	NumberOfLinesToSkip *int64      `json:"numberOfLinesToSkip,omitempty"`
	Provider            string      `json:"provider"`
	RawContent          *string     `json:"rawContent,omitempty"`
	Source              *string     `json:"source,omitempty"`
	SourceType          *SourceType `json:"sourceType,omitempty"`
	TenantId            *string     `json:"tenantId,omitempty"`
	Updated             *string     `json:"updated,omitempty"`
	UpdatedBy           *UserInfo   `json:"updatedBy,omitempty"`
	UploadStatus        *string     `json:"uploadStatus,omitempty"`
	WatchlistAlias      *string     `json:"watchlistAlias,omitempty"`
	WatchlistId         *string     `json:"watchlistId,omitempty"`
	WatchlistType       *string     `json:"watchlistType,omitempty"`
}
