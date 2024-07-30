package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedValue struct {
	Id        *string `json:"id,omitempty"`
	IsActive  *bool   `json:"isActive,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
