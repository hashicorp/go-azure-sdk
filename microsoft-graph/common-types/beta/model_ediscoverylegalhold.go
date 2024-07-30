package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryLegalHold struct {
	ContentQuery         *string                         `json:"contentQuery,omitempty"`
	CreatedBy            *IdentitySet                    `json:"createdBy,omitempty"`
	CreatedDateTime      *string                         `json:"createdDateTime,omitempty"`
	Description          *string                         `json:"description,omitempty"`
	DisplayName          *string                         `json:"displayName,omitempty"`
	Errors               *[]string                       `json:"errors,omitempty"`
	Id                   *string                         `json:"id,omitempty"`
	IsEnabled            *bool                           `json:"isEnabled,omitempty"`
	LastModifiedBy       *IdentitySet                    `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                         `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                         `json:"@odata.type,omitempty"`
	SiteSources          *[]EdiscoverySiteSource         `json:"siteSources,omitempty"`
	Status               *EdiscoveryLegalHoldStatus      `json:"status,omitempty"`
	UnifiedGroupSources  *[]EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`
	UserSources          *[]EdiscoveryUserSource         `json:"userSources,omitempty"`
}
