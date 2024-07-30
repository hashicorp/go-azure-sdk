package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BasicAuthentication struct {
	ODataType *string `json:"@odata.type,omitempty"`
	Password  *string `json:"password,omitempty"`
	Username  *string `json:"username,omitempty"`
}
