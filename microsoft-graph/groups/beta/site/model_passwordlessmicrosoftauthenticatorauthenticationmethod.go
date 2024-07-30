package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordlessMicrosoftAuthenticatorAuthenticationMethod struct {
	CreatedDateTime  *string `json:"createdDateTime,omitempty"`
	CreationDateTime *string `json:"creationDateTime,omitempty"`
	Device           *Device `json:"device,omitempty"`
	DisplayName      *string `json:"displayName,omitempty"`
	Id               *string `json:"id,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
