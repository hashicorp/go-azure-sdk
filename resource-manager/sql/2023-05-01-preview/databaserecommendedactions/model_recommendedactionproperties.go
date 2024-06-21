package databaserecommendedactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionProperties struct {
	Details                    *map[string]interface{}              `json:"details,omitempty"`
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
