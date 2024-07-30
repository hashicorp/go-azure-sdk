package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Bundle struct {
	Album      *Album  `json:"album,omitempty"`
	ChildCount *int64  `json:"childCount,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
}
