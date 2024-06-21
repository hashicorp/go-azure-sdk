package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipal struct {
	AppId               *string `json:"appId,omitempty"`
	CredentialsExpireOn *string `json:"credentialsExpireOn,omitempty"`
	Id                  *string `json:"id,omitempty"`
	TenantId            *string `json:"tenantId,omitempty"`
}
