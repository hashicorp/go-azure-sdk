package sqlpoolsmaintenancewindowoptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceWindowOptionsProperties struct {
	AllowMultipleMaintenanceWindowsPerCycle *bool                         `json:"allowMultipleMaintenanceWindowsPerCycle,omitempty"`
	DefaultDurationInMinutes                *int64                        `json:"defaultDurationInMinutes,omitempty"`
	IsEnabled                               *bool                         `json:"isEnabled,omitempty"`
	MaintenanceWindowCycles                 *[]MaintenanceWindowTimeRange `json:"maintenanceWindowCycles,omitempty"`
	MinCycles                               *int64                        `json:"minCycles,omitempty"`
	MinDurationInMinutes                    *int64                        `json:"minDurationInMinutes,omitempty"`
	TimeGranularityInMinutes                *int64                        `json:"timeGranularityInMinutes,omitempty"`
}
