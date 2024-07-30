package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanDetails struct {
	CategoryDescriptions *PlannerCategoryDescriptions         `json:"categoryDescriptions,omitempty"`
	ContextDetails       *PlannerPlanContextDetailsCollection `json:"contextDetails,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	SharedWith           *PlannerUserIds                      `json:"sharedWith,omitempty"`
}
