package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationGradingCategory struct {
	DisplayName      *string `json:"displayName,omitempty"`
	Id               *string `json:"id,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	PercentageWeight *int64  `json:"percentageWeight,omitempty"`
}
