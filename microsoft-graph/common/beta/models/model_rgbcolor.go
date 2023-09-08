package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RgbColor struct {
	B         *UInt32 `json:"b,omitempty"`
	G         *UInt32 `json:"g,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	R         *UInt32 `json:"r,omitempty"`
}
