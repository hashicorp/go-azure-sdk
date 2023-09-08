package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsCloudPCConnection struct {
	DisplayName           *string `json:"displayName,omitempty"`
	HealthCheckStatus     *string `json:"healthCheckStatus,omitempty"`
	Id                    *string `json:"id,omitempty"`
	LastRefreshedDateTime *string `json:"lastRefreshedDateTime,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	TenantDisplayName     *string `json:"tenantDisplayName,omitempty"`
	TenantId              *string `json:"tenantId,omitempty"`
}
