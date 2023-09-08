package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeSummary struct {
	CompletedDeviceScopeIds        *[]string `json:"completedDeviceScopeIds,omitempty"`
	InsufficientDataDeviceScopeIds *[]string `json:"insufficientDataDeviceScopeIds,omitempty"`
	ODataType                      *string   `json:"@odata.type,omitempty"`
	TotalDeviceScopes              *int64    `json:"totalDeviceScopes,omitempty"`
	TotalDeviceScopesEnabled       *int64    `json:"totalDeviceScopesEnabled,omitempty"`
}
