package sparkjobdefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkServicePlugin struct {
	CleanupStartedAt             *string             `json:"cleanupStartedAt,omitempty"`
	CurrentState                 *PluginCurrentState `json:"currentState,omitempty"`
	MonitoringStartedAt          *string             `json:"monitoringStartedAt,omitempty"`
	PreparationStartedAt         *string             `json:"preparationStartedAt,omitempty"`
	ResourceAcquisitionStartedAt *string             `json:"resourceAcquisitionStartedAt,omitempty"`
	SubmissionStartedAt          *string             `json:"submissionStartedAt,omitempty"`
}

func (o *SparkServicePlugin) GetCleanupStartedAtAsTime() (*time.Time, error) {
	if o.CleanupStartedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CleanupStartedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkServicePlugin) SetCleanupStartedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CleanupStartedAt = &formatted
}

func (o *SparkServicePlugin) GetMonitoringStartedAtAsTime() (*time.Time, error) {
	if o.MonitoringStartedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MonitoringStartedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkServicePlugin) SetMonitoringStartedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MonitoringStartedAt = &formatted
}

func (o *SparkServicePlugin) GetPreparationStartedAtAsTime() (*time.Time, error) {
	if o.PreparationStartedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PreparationStartedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkServicePlugin) SetPreparationStartedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PreparationStartedAt = &formatted
}

func (o *SparkServicePlugin) GetResourceAcquisitionStartedAtAsTime() (*time.Time, error) {
	if o.ResourceAcquisitionStartedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ResourceAcquisitionStartedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkServicePlugin) SetResourceAcquisitionStartedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ResourceAcquisitionStartedAt = &formatted
}

func (o *SparkServicePlugin) GetSubmissionStartedAtAsTime() (*time.Time, error) {
	if o.SubmissionStartedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SubmissionStartedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkServicePlugin) SetSubmissionStartedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SubmissionStartedAt = &formatted
}
