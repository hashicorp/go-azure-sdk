package plannerplan

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainer struct {
	ContainerId *string                   `json:"containerId,omitempty"`
	ODataType   *string                   `json:"@odata.type,omitempty"`
	Type        *PlannerPlanContainerType `json:"type,omitempty"`
	Url         *string                   `json:"url,omitempty"`
}
