package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchEntity struct {
	Acronyms  *[]SearchAcronym  `json:"acronyms,omitempty"`
	Bookmarks *[]SearchBookmark `json:"bookmarks,omitempty"`
	Id        *string           `json:"id,omitempty"`
	ODataType *string           `json:"@odata.type,omitempty"`
	Qnas      *[]SearchQna      `json:"qnas,omitempty"`
}
