package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryHoldPolicy struct {
	ContentQuery         *string                             `json:"contentQuery,omitempty"`
	CreatedBy            *IdentitySet                        `json:"createdBy,omitempty"`
	CreatedDateTime      *string                             `json:"createdDateTime,omitempty"`
	Description          *string                             `json:"description,omitempty"`
	DisplayName          *string                             `json:"displayName,omitempty"`
	Errors               *[]string                           `json:"errors,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
	IsEnabled            *bool                               `json:"isEnabled,omitempty"`
	LastModifiedBy       *IdentitySet                        `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                             `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                             `json:"@odata.type,omitempty"`
	SiteSources          *[]SecuritySiteSource               `json:"siteSources,omitempty"`
	Status               *SecurityEdiscoveryHoldPolicyStatus `json:"status,omitempty"`
	UserSources          *[]SecurityUserSource               `json:"userSources,omitempty"`
}
