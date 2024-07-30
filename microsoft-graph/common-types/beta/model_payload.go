package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Payload struct {
	Brand                   *PayloadBrand                `json:"brand,omitempty"`
	Complexity              *PayloadComplexity           `json:"complexity,omitempty"`
	CreatedBy               *EmailIdentity               `json:"createdBy,omitempty"`
	CreatedDateTime         *string                      `json:"createdDateTime,omitempty"`
	Description             *string                      `json:"description,omitempty"`
	Detail                  *PayloadDetail               `json:"detail,omitempty"`
	DisplayName             *string                      `json:"displayName,omitempty"`
	Id                      *string                      `json:"id,omitempty"`
	Industry                *PayloadIndustry             `json:"industry,omitempty"`
	IsAutomated             *bool                        `json:"isAutomated,omitempty"`
	IsControversial         *bool                        `json:"isControversial,omitempty"`
	IsCurrentEvent          *bool                        `json:"isCurrentEvent,omitempty"`
	Language                *string                      `json:"language,omitempty"`
	LastModifiedBy          *EmailIdentity               `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime    *string                      `json:"lastModifiedDateTime,omitempty"`
	ODataType               *string                      `json:"@odata.type,omitempty"`
	PayloadTags             *[]string                    `json:"payloadTags,omitempty"`
	Platform                *PayloadPlatform             `json:"platform,omitempty"`
	PredictedCompromiseRate *float64                     `json:"predictedCompromiseRate,omitempty"`
	SimulationAttackType    *PayloadSimulationAttackType `json:"simulationAttackType,omitempty"`
	Source                  *PayloadSource               `json:"source,omitempty"`
	Status                  *PayloadStatus               `json:"status,omitempty"`
	Technique               *PayloadTechnique            `json:"technique,omitempty"`
	Theme                   *PayloadTheme                `json:"theme,omitempty"`
}
