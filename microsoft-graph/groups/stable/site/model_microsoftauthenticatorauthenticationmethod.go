package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethod struct {
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	Device          *Device `json:"device,omitempty"`
	DeviceTag       *string `json:"deviceTag,omitempty"`
	DisplayName     *string `json:"displayName,omitempty"`
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	PhoneAppVersion *string `json:"phoneAppVersion,omitempty"`
}
