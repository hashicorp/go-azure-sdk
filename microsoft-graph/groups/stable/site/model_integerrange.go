package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegerRange struct {
	End       *int64  `json:"end,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Start     *int64  `json:"start,omitempty"`
}
