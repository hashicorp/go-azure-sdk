package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuth2PermissionGrant struct {
	ClientId    *string `json:"clientId,omitempty"`
	ConsentType *string `json:"consentType,omitempty"`
	ExpiryTime  *string `json:"expiryTime,omitempty"`
	Id          *string `json:"id,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
	ResourceId  *string `json:"resourceId,omitempty"`
	Scope       *string `json:"scope,omitempty"`
	StartTime   *string `json:"startTime,omitempty"`
}
