package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantCustomizedInformation struct {
	Contacts    *[]ManagedTenantsTenantContactInformation `json:"contacts,omitempty"`
	DisplayName *string                                   `json:"displayName,omitempty"`
	Id          *string                                   `json:"id,omitempty"`
	ODataType   *string                                   `json:"@odata.type,omitempty"`
	TenantId    *string                                   `json:"tenantId,omitempty"`
	Website     *string                                   `json:"website,omitempty"`
}
