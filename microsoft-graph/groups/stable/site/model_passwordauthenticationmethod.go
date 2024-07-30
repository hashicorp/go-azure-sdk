package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordAuthenticationMethod struct {
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	Password        *string `json:"password,omitempty"`
}
