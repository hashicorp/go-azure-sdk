package managedmaintenancewindowstatus

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *ManagedMaintenanceWindowStatus) GetLastWindowEndTimeUTCAsTime() (*time.Time, error) {
	if o.LastWindowEndTimeUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastWindowEndTimeUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedMaintenanceWindowStatus) SetLastWindowEndTimeUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastWindowEndTimeUTC = &formatted
}

func (o *ManagedMaintenanceWindowStatus) GetLastWindowStartTimeUTCAsTime() (*time.Time, error) {
	if o.LastWindowStartTimeUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastWindowStartTimeUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedMaintenanceWindowStatus) SetLastWindowStartTimeUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastWindowStartTimeUTC = &formatted
}

func (o *ManagedMaintenanceWindowStatus) GetLastWindowStatusUpdateAtUTCAsTime() (*time.Time, error) {
	if o.LastWindowStatusUpdateAtUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastWindowStatusUpdateAtUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedMaintenanceWindowStatus) SetLastWindowStatusUpdateAtUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastWindowStatusUpdateAtUTC = &formatted
}
