package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScope struct {
	CreatedDateTime      *string                                      `json:"createdDateTime,omitempty"`
	DeviceScopeName      *string                                      `json:"deviceScopeName,omitempty"`
	Enabled              *bool                                        `json:"enabled,omitempty"`
	Id                   *string                                      `json:"id,omitempty"`
	IsBuiltIn            *bool                                        `json:"isBuiltIn,omitempty"`
	LastModifiedDateTime *string                                      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	Operator             *UserExperienceAnalyticsDeviceScopeOperator  `json:"operator,omitempty"`
	OwnerId              *string                                      `json:"ownerId,omitempty"`
	Parameter            *UserExperienceAnalyticsDeviceScopeParameter `json:"parameter,omitempty"`
	Status               *UserExperienceAnalyticsDeviceScopeStatus    `json:"status,omitempty"`
	Value                *string                                      `json:"value,omitempty"`
	ValueObjectId        *string                                      `json:"valueObjectId,omitempty"`
}
