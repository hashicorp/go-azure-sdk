package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsApplication struct {
	ODataType    *string   `json:"@odata.type,omitempty"`
	PackageSid   *string   `json:"packageSid,omitempty"`
	RedirectUris *[]string `json:"redirectUris,omitempty"`
}
