package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCountMetric struct {
	Count     *int64  `json:"count,omitempty"`
	FactDate  *string `json:"factDate,omitempty"`
	Id        *string `json:"id,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
