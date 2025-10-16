package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalInformation struct {
	ApplicationId string `json:"applicationId"`
	Password      string `json:"password"`
	PrincipalId   string `json:"principalId"`
	TenantId      string `json:"tenantId"`
}
