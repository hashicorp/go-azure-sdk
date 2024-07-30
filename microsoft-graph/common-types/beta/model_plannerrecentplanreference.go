package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRecentPlanReference struct {
	LastAccessedDateTime *string `json:"lastAccessedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	PlanTitle            *string `json:"planTitle,omitempty"`
}
