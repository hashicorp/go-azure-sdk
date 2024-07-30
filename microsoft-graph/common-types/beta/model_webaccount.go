package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAccount struct {
	AllowedAudiences     *WebAccountAllowedAudiences `json:"allowedAudiences,omitempty"`
	CreatedBy            *IdentitySet                `json:"createdBy,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Inference            *InferenceData              `json:"inference,omitempty"`
	IsSearchable         *bool                       `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Service              *ServiceInformation         `json:"service,omitempty"`
	Source               *PersonDataSources          `json:"source,omitempty"`
	StatusMessage        *string                     `json:"statusMessage,omitempty"`
	ThumbnailUrl         *string                     `json:"thumbnailUrl,omitempty"`
	UserId               *string                     `json:"userId,omitempty"`
	WebUrl               *string                     `json:"webUrl,omitempty"`
}
