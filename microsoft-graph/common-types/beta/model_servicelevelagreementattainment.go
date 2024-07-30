package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceLevelAgreementAttainment struct {
	EndDate   *string  `json:"endDate,omitempty"`
	ODataType *string  `json:"@odata.type,omitempty"`
	Score     *float64 `json:"score,omitempty"`
	StartDate *string  `json:"startDate,omitempty"`
}
