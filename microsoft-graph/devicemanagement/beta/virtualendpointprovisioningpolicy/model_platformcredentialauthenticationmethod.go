package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformCredentialAuthenticationMethod struct {
	CreatedDateTime *string                                            `json:"createdDateTime,omitempty"`
	Device          *Device                                            `json:"device,omitempty"`
	DisplayName     *string                                            `json:"displayName,omitempty"`
	Id              *string                                            `json:"id,omitempty"`
	KeyStrength     *PlatformCredentialAuthenticationMethodKeyStrength `json:"keyStrength,omitempty"`
	ODataType       *string                                            `json:"@odata.type,omitempty"`
	Platform        *PlatformCredentialAuthenticationMethodPlatform    `json:"platform,omitempty"`
}
