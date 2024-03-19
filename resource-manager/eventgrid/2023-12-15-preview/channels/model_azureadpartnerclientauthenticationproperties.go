package channels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADPartnerClientAuthenticationProperties struct {
	AzureActiveDirectoryApplicationIdOrUri *string `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`
	AzureActiveDirectoryTenantId           *string `json:"azureActiveDirectoryTenantId,omitempty"`
}
