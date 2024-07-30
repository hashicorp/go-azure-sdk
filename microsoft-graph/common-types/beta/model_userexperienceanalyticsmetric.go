package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsMetric struct {
	Id        *string  `json:"id,omitempty"`
	ODataType *string  `json:"@odata.type,omitempty"`
	Unit      *string  `json:"unit,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}
