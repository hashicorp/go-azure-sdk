package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessWebCategoriesSummary struct {
	DeviceCount      *int64                    `json:"deviceCount,omitempty"`
	ODataType        *string                   `json:"@odata.type,omitempty"`
	TransactionCount *int64                    `json:"transactionCount,omitempty"`
	UserCount        *int64                    `json:"userCount,omitempty"`
	WebCategory      *NetworkaccessWebCategory `json:"webCategory,omitempty"`
}
