package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTypeOrder struct {
	Default   *bool   `json:"default,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Position  *int64  `json:"position,omitempty"`
}
