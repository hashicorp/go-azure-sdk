package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunRoleCountMetric struct {
	Count     *int64  `json:"count,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Role      *string `json:"role,omitempty"`
}
