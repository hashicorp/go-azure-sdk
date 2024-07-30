package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantContactInformation struct {
	Email     *string `json:"email,omitempty"`
	Name      *string `json:"name,omitempty"`
	Notes     *string `json:"notes,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Title     *string `json:"title,omitempty"`
}
