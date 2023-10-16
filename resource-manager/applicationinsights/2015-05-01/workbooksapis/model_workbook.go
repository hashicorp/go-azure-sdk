package workbooksapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Workbook struct {
	Id         *string             `json:"id,omitempty"`
	Kind       *SharedTypeKind     `json:"kind,omitempty"`
	Location   *string             `json:"location,omitempty"`
	Name       *string             `json:"name,omitempty"`
	Properties *WorkbookProperties `json:"properties,omitempty"`
	Tags       *map[string]string  `json:"tags,omitempty"`
	Type       *string             `json:"type,omitempty"`
}
