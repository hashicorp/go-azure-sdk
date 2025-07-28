package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadIdentityFederation struct {
	AppId    *string `json:"appId,omitempty"`
	Id       *string `json:"id,omitempty"`
	Issuer   *string `json:"issuer,omitempty"`
	Subject  *string `json:"subject,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
}
