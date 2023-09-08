package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationReportOverview struct {
	ODataType               *string                  `json:"@odata.type,omitempty"`
	RecommendedActions      *[]RecommendedAction     `json:"recommendedActions,omitempty"`
	ResolvedTargetsCount    *int64                   `json:"resolvedTargetsCount,omitempty"`
	SimulationEventsContent *SimulationEventsContent `json:"simulationEventsContent,omitempty"`
	TrainingEventsContent   *TrainingEventsContent   `json:"trainingEventsContent,omitempty"`
}
