package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCustodian struct {
	AcknowledgedDateTime *string                         `json:"acknowledgedDateTime,omitempty"`
	ApplyHoldToSources   *bool                           `json:"applyHoldToSources,omitempty"`
	CreatedDateTime      *string                         `json:"createdDateTime,omitempty"`
	DisplayName          *string                         `json:"displayName,omitempty"`
	Email                *string                         `json:"email,omitempty"`
	HoldStatus           *EdiscoveryCustodianHoldStatus  `json:"holdStatus,omitempty"`
	Id                   *string                         `json:"id,omitempty"`
	LastIndexOperation   *EdiscoveryCaseIndexOperation   `json:"lastIndexOperation,omitempty"`
	LastModifiedDateTime *string                         `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                         `json:"@odata.type,omitempty"`
	ReleasedDateTime     *string                         `json:"releasedDateTime,omitempty"`
	SiteSources          *[]EdiscoverySiteSource         `json:"siteSources,omitempty"`
	Status               *EdiscoveryCustodianStatus      `json:"status,omitempty"`
	UnifiedGroupSources  *[]EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`
	UserSources          *[]EdiscoveryUserSource         `json:"userSources,omitempty"`
}
