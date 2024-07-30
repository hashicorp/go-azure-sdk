package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerUser struct {
	Id        *string        `json:"id,omitempty"`
	ODataType *string        `json:"@odata.type,omitempty"`
	Plans     *[]PlannerPlan `json:"plans,omitempty"`
	Tasks     *[]PlannerTask `json:"tasks,omitempty"`
}
