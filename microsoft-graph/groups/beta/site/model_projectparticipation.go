package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectParticipation struct {
	AllowedAudiences     *ProjectParticipationAllowedAudiences `json:"allowedAudiences,omitempty"`
	Categories           *[]string                             `json:"categories,omitempty"`
	Client               *CompanyDetail                        `json:"client,omitempty"`
	CollaborationTags    *[]string                             `json:"collaborationTags,omitempty"`
	Colleagues           *[]RelatedPerson                      `json:"colleagues,omitempty"`
	CreatedBy            *IdentitySet                          `json:"createdBy,omitempty"`
	CreatedDateTime      *string                               `json:"createdDateTime,omitempty"`
	Detail               *PositionDetail                       `json:"detail,omitempty"`
	DisplayName          *string                               `json:"displayName,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	Inference            *InferenceData                        `json:"inference,omitempty"`
	IsSearchable         *bool                                 `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                          `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                               `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	Source               *PersonDataSources                    `json:"source,omitempty"`
	Sponsors             *[]RelatedPerson                      `json:"sponsors,omitempty"`
	ThumbnailUrl         *string                               `json:"thumbnailUrl,omitempty"`
}
