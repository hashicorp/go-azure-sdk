package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantInfo struct {
	ODataType *string `json:"@odata.type,omitempty"`
	TenantId  *string `json:"tenantId,omitempty"`
}
