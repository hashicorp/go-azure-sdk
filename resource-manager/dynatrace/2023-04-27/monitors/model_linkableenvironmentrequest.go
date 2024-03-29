package monitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkableEnvironmentRequest struct {
	Region        string `json:"region"`
	TenantId      string `json:"tenantId"`
	UserPrincipal string `json:"userPrincipal"`
}
