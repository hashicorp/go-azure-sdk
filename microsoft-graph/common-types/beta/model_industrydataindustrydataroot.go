package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRoot struct {
	DataConnectors       *[]IndustryDataIndustryDataConnector    `json:"dataConnectors,omitempty"`
	Id                   *string                                 `json:"id,omitempty"`
	InboundFlows         *[]IndustryDataInboundFlow              `json:"inboundFlows,omitempty"`
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	Operations           *[]LongRunningOperation                 `json:"operations,omitempty"`
	ReferenceDefinitions *[]IndustryDataReferenceDefinition      `json:"referenceDefinitions,omitempty"`
	RoleGroups           *[]IndustryDataRoleGroup                `json:"roleGroups,omitempty"`
	Runs                 *[]IndustryDataIndustryDataRun          `json:"runs,omitempty"`
	SourceSystems        *[]IndustryDataSourceSystemDefinition   `json:"sourceSystems,omitempty"`
	Years                *[]IndustryDataYearTimePeriodDefinition `json:"years,omitempty"`
}
