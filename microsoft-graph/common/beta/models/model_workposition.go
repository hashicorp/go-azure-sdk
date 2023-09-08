package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkPosition struct {
	AllowedAudiences     *WorkPositionAllowedAudiences `json:"allowedAudiences,omitempty"`
	Categories           *[]string                     `json:"categories,omitempty"`
	Colleagues           *[]RelatedPerson              `json:"colleagues,omitempty"`
	CreatedBy            *IdentitySet                  `json:"createdBy,omitempty"`
	CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
	Detail               *PositionDetail               `json:"detail,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	Inference            *InferenceData                `json:"inference,omitempty"`
	IsCurrent            *bool                         `json:"isCurrent,omitempty"`
	IsSearchable         *bool                         `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                  `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                       `json:"lastModifiedDateTime,omitempty"`
	Manager              *RelatedPerson                `json:"manager,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	Source               *PersonDataSources            `json:"source,omitempty"`
}
