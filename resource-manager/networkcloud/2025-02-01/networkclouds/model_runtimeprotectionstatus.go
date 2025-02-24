package networkclouds

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuntimeProtectionStatus struct {
	DefinitionsLastUpdated *string `json:"definitionsLastUpdated,omitempty"`
	DefinitionsVersion     *string `json:"definitionsVersion,omitempty"`
	ScanCompletedTime      *string `json:"scanCompletedTime,omitempty"`
	ScanScheduledTime      *string `json:"scanScheduledTime,omitempty"`
	ScanStartedTime        *string `json:"scanStartedTime,omitempty"`
}

func (o *RuntimeProtectionStatus) GetDefinitionsLastUpdatedAsTime() (*time.Time, error) {
	if o.DefinitionsLastUpdated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DefinitionsLastUpdated, "2006-01-02T15:04:05Z07:00")
}

func (o *RuntimeProtectionStatus) SetDefinitionsLastUpdatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DefinitionsLastUpdated = &formatted
}

func (o *RuntimeProtectionStatus) GetScanCompletedTimeAsTime() (*time.Time, error) {
	if o.ScanCompletedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScanCompletedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RuntimeProtectionStatus) SetScanCompletedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScanCompletedTime = &formatted
}

func (o *RuntimeProtectionStatus) GetScanScheduledTimeAsTime() (*time.Time, error) {
	if o.ScanScheduledTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScanScheduledTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RuntimeProtectionStatus) SetScanScheduledTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScanScheduledTime = &formatted
}

func (o *RuntimeProtectionStatus) GetScanStartedTimeAsTime() (*time.Time, error) {
	if o.ScanStartedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScanStartedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RuntimeProtectionStatus) SetScanStartedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScanStartedTime = &formatted
}
