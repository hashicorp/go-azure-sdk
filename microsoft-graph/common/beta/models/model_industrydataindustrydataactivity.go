package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivity struct {
	DisplayName     *string                                          `json:"displayName,omitempty"`
	Id              *string                                          `json:"id,omitempty"`
	ODataType       *string                                          `json:"@odata.type,omitempty"`
	ReadinessStatus *IndustryDataIndustryDataActivityReadinessStatus `json:"readinessStatus,omitempty"`
}
