package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFileFlow struct {
	DataConnector      *IndustryDataIndustryDataConnector          `json:"dataConnector,omitempty"`
	DataDomain         *IndustryDataInboundFileFlowDataDomain      `json:"dataDomain,omitempty"`
	DisplayName        *string                                     `json:"displayName,omitempty"`
	EffectiveDateTime  *string                                     `json:"effectiveDateTime,omitempty"`
	ExpirationDateTime *string                                     `json:"expirationDateTime,omitempty"`
	Id                 *string                                     `json:"id,omitempty"`
	ODataType          *string                                     `json:"@odata.type,omitempty"`
	ReadinessStatus    *IndustryDataInboundFileFlowReadinessStatus `json:"readinessStatus,omitempty"`
	Year               *IndustryDataYearTimePeriodDefinition       `json:"year,omitempty"`
}
