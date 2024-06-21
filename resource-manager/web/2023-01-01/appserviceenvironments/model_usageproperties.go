package appserviceenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageProperties struct {
	ComputeMode   *ComputeModeOptions `json:"computeMode,omitempty"`
	CurrentValue  *int64              `json:"currentValue,omitempty"`
	DisplayName   *string             `json:"displayName,omitempty"`
	Limit         *int64              `json:"limit,omitempty"`
	NextResetTime *string             `json:"nextResetTime,omitempty"`
	ResourceName  *string             `json:"resourceName,omitempty"`
	SiteMode      *string             `json:"siteMode,omitempty"`
	Unit          *string             `json:"unit,omitempty"`
}
