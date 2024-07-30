package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSecureScore struct {
	CreateDateTime *string `json:"createDateTime,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	TenantMaxScore *int64  `json:"tenantMaxScore,omitempty"`
	TenantScore    *int64  `json:"tenantScore,omitempty"`
}
