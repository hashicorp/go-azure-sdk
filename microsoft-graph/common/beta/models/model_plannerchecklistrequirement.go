package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerChecklistRequirement struct {
	ODataType                *string   `json:"@odata.type,omitempty"`
	RequiredChecklistItemIds *[]string `json:"requiredChecklistItemIds,omitempty"`
}
