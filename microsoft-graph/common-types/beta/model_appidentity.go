package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppIdentity struct {
	AppId                *string `json:"appId,omitempty"`
	DisplayName          *string `json:"displayName,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	ServicePrincipalId   *string `json:"servicePrincipalId,omitempty"`
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
}
