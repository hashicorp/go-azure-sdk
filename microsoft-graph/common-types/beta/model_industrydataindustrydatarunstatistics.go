package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunStatistics struct {
	ActivityStatistics *[]IndustryDataIndustryDataActivityStatistics `json:"activityStatistics,omitempty"`
	InboundTotals      *IndustryDataAggregatedInboundStatistics      `json:"inboundTotals,omitempty"`
	ODataType          *string                                       `json:"@odata.type,omitempty"`
	RunId              *string                                       `json:"runId,omitempty"`
	Status             *IndustryDataIndustryDataRunStatisticsStatus  `json:"status,omitempty"`
}
