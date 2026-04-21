package databaseadvisors

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionProperties struct {
	Details                    *map[string]string                   `json:"details,omitempty"`
	ErrorDetails               *RecommendedActionErrorInfo          `json:"errorDetails,omitempty"`
	EstimatedImpact            *[]RecommendedActionImpactRecord     `json:"estimatedImpact,omitempty"`
	ExecuteActionDuration      *string                              `json:"executeActionDuration,omitempty"`
	ExecuteActionInitiatedBy   *RecommendedActionInitiatedBy        `json:"executeActionInitiatedBy,omitempty"`
	ExecuteActionInitiatedTime *string                              `json:"executeActionInitiatedTime,omitempty"`
	ExecuteActionStartTime     *string                              `json:"executeActionStartTime,omitempty"`
	ImplementationDetails      *RecommendedActionImplementationInfo `json:"implementationDetails,omitempty"`
	IsArchivedAction           *bool                                `json:"isArchivedAction,omitempty"`
	IsExecutableAction         *bool                                `json:"isExecutableAction,omitempty"`
	IsRevertableAction         *bool                                `json:"isRevertableAction,omitempty"`
	LastRefresh                *string                              `json:"lastRefresh,omitempty"`
	LinkedObjects              *[]string                            `json:"linkedObjects,omitempty"`
	ObservedImpact             *[]RecommendedActionImpactRecord     `json:"observedImpact,omitempty"`
	RecommendationReason       *string                              `json:"recommendationReason,omitempty"`
	RevertActionDuration       *string                              `json:"revertActionDuration,omitempty"`
	RevertActionInitiatedBy    *RecommendedActionInitiatedBy        `json:"revertActionInitiatedBy,omitempty"`
	RevertActionInitiatedTime  *string                              `json:"revertActionInitiatedTime,omitempty"`
	RevertActionStartTime      *string                              `json:"revertActionStartTime,omitempty"`
	Score                      *int64                               `json:"score,omitempty"`
	State                      RecommendedActionStateInfo           `json:"state"`
	TimeSeries                 *[]RecommendedActionMetricInfo       `json:"timeSeries,omitempty"`
	ValidSince                 *string                              `json:"validSince,omitempty"`
}

func (o *RecommendedActionProperties) GetExecuteActionInitiatedTimeAsTime() (*time.Time, error) {
	if o.ExecuteActionInitiatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExecuteActionInitiatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionProperties) SetExecuteActionInitiatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExecuteActionInitiatedTime = &formatted
}

func (o *RecommendedActionProperties) GetExecuteActionStartTimeAsTime() (*time.Time, error) {
	if o.ExecuteActionStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExecuteActionStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionProperties) SetExecuteActionStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExecuteActionStartTime = &formatted
}

func (o *RecommendedActionProperties) GetLastRefreshAsTime() (*time.Time, error) {
	if o.LastRefresh == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRefresh, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionProperties) SetLastRefreshAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRefresh = &formatted
}

func (o *RecommendedActionProperties) GetRevertActionInitiatedTimeAsTime() (*time.Time, error) {
	if o.RevertActionInitiatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RevertActionInitiatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionProperties) SetRevertActionInitiatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RevertActionInitiatedTime = &formatted
}

func (o *RecommendedActionProperties) GetRevertActionStartTimeAsTime() (*time.Time, error) {
	if o.RevertActionStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RevertActionStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionProperties) SetRevertActionStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RevertActionStartTime = &formatted
}

func (o *RecommendedActionProperties) GetValidSinceAsTime() (*time.Time, error) {
	if o.ValidSince == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidSince, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionProperties) SetValidSinceAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidSince = &formatted
}
