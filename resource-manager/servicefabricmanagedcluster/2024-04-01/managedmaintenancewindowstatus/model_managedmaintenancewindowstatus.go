package managedmaintenancewindowstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMaintenanceWindowStatus struct {
	CanApplyUpdates             *bool   `json:"canApplyUpdates,omitempty"`
	IsRegionReady               *bool   `json:"isRegionReady,omitempty"`
	IsWindowActive              *bool   `json:"isWindowActive,omitempty"`
	IsWindowEnabled             *bool   `json:"isWindowEnabled,omitempty"`
	LastWindowEndTimeUTC        *string `json:"lastWindowEndTimeUTC,omitempty"`
	LastWindowStartTimeUTC      *string `json:"lastWindowStartTimeUTC,omitempty"`
	LastWindowStatusUpdateAtUTC *string `json:"lastWindowStatusUpdateAtUTC,omitempty"`
}
