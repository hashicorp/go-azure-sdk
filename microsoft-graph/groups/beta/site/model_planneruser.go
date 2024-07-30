package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerUser struct {
	All                    *[]PlannerDelta                         `json:"all,omitempty"`
	FavoritePlanReferences *PlannerFavoritePlanReferenceCollection `json:"favoritePlanReferences,omitempty"`
	FavoritePlans          *[]PlannerPlan                          `json:"favoritePlans,omitempty"`
	Id                     *string                                 `json:"id,omitempty"`
	ODataType              *string                                 `json:"@odata.type,omitempty"`
	Plans                  *[]PlannerPlan                          `json:"plans,omitempty"`
	RecentPlanReferences   *PlannerRecentPlanReferenceCollection   `json:"recentPlanReferences,omitempty"`
	RecentPlans            *[]PlannerPlan                          `json:"recentPlans,omitempty"`
	RosterPlans            *[]PlannerPlan                          `json:"rosterPlans,omitempty"`
	Tasks                  *[]PlannerTask                          `json:"tasks,omitempty"`
}
