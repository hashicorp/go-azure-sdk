package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosBookmark struct {
	BookmarkFolder *string `json:"bookmarkFolder,omitempty"`
	DisplayName    *string `json:"displayName,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	Url            *string `json:"url,omitempty"`
}
