package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Document struct {
	Comments  *[]DocumentComment `json:"comments,omitempty"`
	Id        *string            `json:"id,omitempty"`
	ODataType *string            `json:"@odata.type,omitempty"`
}
