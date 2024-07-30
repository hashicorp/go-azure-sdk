package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightsRoot struct {
	Daily     *DailyUserInsightMetricsRoot   `json:"daily,omitempty"`
	Id        *string                        `json:"id,omitempty"`
	Monthly   *MonthlyUserInsightMetricsRoot `json:"monthly,omitempty"`
	ODataType *string                        `json:"@odata.type,omitempty"`
}
