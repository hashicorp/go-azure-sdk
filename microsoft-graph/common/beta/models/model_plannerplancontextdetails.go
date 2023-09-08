package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextDetails struct {
	CustomLinkText  *string                                   `json:"customLinkText,omitempty"`
	DisplayLinkType *PlannerPlanContextDetailsDisplayLinkType `json:"displayLinkType,omitempty"`
	ODataType       *string                                   `json:"@odata.type,omitempty"`
	State           *PlannerPlanContextDetailsState           `json:"state,omitempty"`
	Url             *string                                   `json:"url,omitempty"`
}
