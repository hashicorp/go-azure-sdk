package post

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureActiveDirectoryApp struct {
	AppKey                 *string `json:"appKey,omitempty"`
	ApplicationId          *string `json:"applicationId,omitempty"`
	IgnoreAzurePermissions *bool   `json:"ignoreAzurePermissions,omitempty"`
	TenantId               *string `json:"tenantId,omitempty"`
}
