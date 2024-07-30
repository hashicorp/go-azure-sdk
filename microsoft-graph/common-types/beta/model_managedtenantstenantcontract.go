package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantContract struct {
	ContractType      *int64  `json:"contractType,omitempty"`
	DefaultDomainName *string `json:"defaultDomainName,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
}
