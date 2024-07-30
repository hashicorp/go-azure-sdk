package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlow struct {
	DataConnector      *IndustryDataIndustryDataConnector      `json:"dataConnector,omitempty"`
	DataDomain         *IndustryDataInboundFlowDataDomain      `json:"dataDomain,omitempty"`
	DisplayName        *string                                 `json:"displayName,omitempty"`
	EffectiveDateTime  *string                                 `json:"effectiveDateTime,omitempty"`
	ExpirationDateTime *string                                 `json:"expirationDateTime,omitempty"`
	Id                 *string                                 `json:"id,omitempty"`
	ODataType          *string                                 `json:"@odata.type,omitempty"`
	ReadinessStatus    *IndustryDataInboundFlowReadinessStatus `json:"readinessStatus,omitempty"`
	Year               *IndustryDataYearTimePeriodDefinition   `json:"year,omitempty"`
}
