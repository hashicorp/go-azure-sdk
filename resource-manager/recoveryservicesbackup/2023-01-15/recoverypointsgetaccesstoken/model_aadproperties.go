package recoverypointsgetaccesstoken

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AADProperties struct {
	Audience                 *string `json:"audience,omitempty"`
	Authority                *string `json:"authority,omitempty"`
	ServicePrincipalClientId *string `json:"servicePrincipalClientId,omitempty"`
	ServicePrincipalObjectId *string `json:"servicePrincipalObjectId,omitempty"`
	TenantId                 *string `json:"tenantId,omitempty"`
}
