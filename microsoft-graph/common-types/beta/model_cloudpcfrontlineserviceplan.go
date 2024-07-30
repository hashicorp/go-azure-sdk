package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCFrontLineServicePlan struct {
	DisplayName *string `json:"displayName,omitempty"`
	Id          *string `json:"id,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	TotalCount  *int64  `json:"totalCount,omitempty"`
	UsedCount   *int64  `json:"usedCount,omitempty"`
}
