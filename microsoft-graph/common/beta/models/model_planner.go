package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Planner struct {
	Buckets   *[]PlannerBucket `json:"buckets,omitempty"`
	ODataType *string          `json:"@odata.type,omitempty"`
	Plans     *[]PlannerPlan   `json:"plans,omitempty"`
	Rosters   *[]PlannerRoster `json:"rosters,omitempty"`
	Tasks     *[]PlannerTask   `json:"tasks,omitempty"`
}
