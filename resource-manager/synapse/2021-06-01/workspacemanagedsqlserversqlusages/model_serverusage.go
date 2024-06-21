package workspacemanagedsqlserversqlusages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerUsage struct {
	CurrentValue  *float64 `json:"currentValue,omitempty"`
	DisplayName   *string  `json:"displayName,omitempty"`
	Limit         *float64 `json:"limit,omitempty"`
	Name          *string  `json:"name,omitempty"`
	NextResetTime *string  `json:"nextResetTime,omitempty"`
	ResourceName  *string  `json:"resourceName,omitempty"`
	Unit          *string  `json:"unit,omitempty"`
}
