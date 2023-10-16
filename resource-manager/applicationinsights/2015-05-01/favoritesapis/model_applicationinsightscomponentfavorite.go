package favoritesapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationInsightsComponentFavorite struct {
	Category                *string       `json:"Category,omitempty"`
	Config                  *string       `json:"Config,omitempty"`
	FavoriteId              *string       `json:"FavoriteId,omitempty"`
	FavoriteType            *FavoriteType `json:"FavoriteType,omitempty"`
	IsGeneratedFromTemplate *bool         `json:"IsGeneratedFromTemplate,omitempty"`
	Name                    *string       `json:"Name,omitempty"`
	SourceType              *string       `json:"SourceType,omitempty"`
	Tags                    *[]string     `json:"Tags,omitempty"`
	TimeModified            *string       `json:"TimeModified,omitempty"`
	UserId                  *string       `json:"UserId,omitempty"`
	Version                 *string       `json:"Version,omitempty"`
}
