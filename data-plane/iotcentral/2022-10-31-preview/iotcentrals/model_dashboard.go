package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Dashboard struct {
	DisplayName   string    `json:"displayName"`
	Etag          *string   `json:"etag,omitempty"`
	Favorite      *bool     `json:"favorite,omitempty"`
	Id            *string   `json:"id,omitempty"`
	Organizations *[]string `json:"organizations,omitempty"`
	Personal      *bool     `json:"personal,omitempty"`
	Tiles         *[]Tile   `json:"tiles,omitempty"`
}
