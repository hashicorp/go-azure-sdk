package tenants

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantIdDescription struct {
	Id       *string `json:"id,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
}
