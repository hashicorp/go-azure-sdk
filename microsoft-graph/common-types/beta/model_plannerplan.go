package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlan struct {
	Buckets              *[]PlannerBucket              `json:"buckets,omitempty"`
	Container            *PlannerPlanContainer         `json:"container,omitempty"`
	Contexts             *PlannerPlanContextCollection `json:"contexts,omitempty"`
	CreatedBy            *IdentitySet                  `json:"createdBy,omitempty"`
	CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
	CreationSource       *PlannerPlanCreation          `json:"creationSource,omitempty"`
	Details              *PlannerPlanDetails           `json:"details,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	Owner                *string                       `json:"owner,omitempty"`
	SharedWithContainers *[]PlannerSharedWithContainer `json:"sharedWithContainers,omitempty"`
	Tasks                *[]PlannerTask                `json:"tasks,omitempty"`
	Title                *string                       `json:"title,omitempty"`
}
