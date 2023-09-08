package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskLocalUser struct {
	ODataType *string `json:"@odata.type,omitempty"`
	UserName  *string `json:"userName,omitempty"`
}
