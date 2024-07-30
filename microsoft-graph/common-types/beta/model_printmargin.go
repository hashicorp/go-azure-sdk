package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintMargin struct {
	Bottom    *int64  `json:"bottom,omitempty"`
	Left      *int64  `json:"left,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Right     *int64  `json:"right,omitempty"`
	Top       *int64  `json:"top,omitempty"`
}
