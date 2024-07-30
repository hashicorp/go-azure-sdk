package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHelloForBusinessAuthenticationMethod struct {
	CreatedDateTime *string                                                 `json:"createdDateTime,omitempty"`
	Device          *Device                                                 `json:"device,omitempty"`
	DisplayName     *string                                                 `json:"displayName,omitempty"`
	Id              *string                                                 `json:"id,omitempty"`
	KeyStrength     *WindowsHelloForBusinessAuthenticationMethodKeyStrength `json:"keyStrength,omitempty"`
	ODataType       *string                                                 `json:"@odata.type,omitempty"`
}
