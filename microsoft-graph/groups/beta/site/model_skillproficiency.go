package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkillProficiency struct {
	AllowedAudiences     *SkillProficiencyAllowedAudiences `json:"allowedAudiences,omitempty"`
	Categories           *[]string                         `json:"categories,omitempty"`
	CollaborationTags    *[]string                         `json:"collaborationTags,omitempty"`
	CreatedBy            *IdentitySet                      `json:"createdBy,omitempty"`
	CreatedDateTime      *string                           `json:"createdDateTime,omitempty"`
	DisplayName          *string                           `json:"displayName,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	Inference            *InferenceData                    `json:"inference,omitempty"`
	IsSearchable         *bool                             `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                      `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                           `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                           `json:"@odata.type,omitempty"`
	Proficiency          *SkillProficiencyProficiency      `json:"proficiency,omitempty"`
	Source               *PersonDataSources                `json:"source,omitempty"`
	ThumbnailUrl         *string                           `json:"thumbnailUrl,omitempty"`
	WebUrl               *string                           `json:"webUrl,omitempty"`
}
