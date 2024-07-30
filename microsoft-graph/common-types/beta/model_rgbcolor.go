package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RgbColor struct {
	B         *int64  `json:"b,omitempty"`
	G         *int64  `json:"g,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	R         *int64  `json:"r,omitempty"`
}
