package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAuthenticationMethod struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
	Id           *string `json:"id,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
}
