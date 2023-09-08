package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenario struct {
	CreatedBy            *IdentitySet             `json:"createdBy,omitempty"`
	CreatedDateTime      *string                  `json:"createdDateTime,omitempty"`
	DisplayName          *string                  `json:"displayName,omitempty"`
	Id                   *string                  `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet             `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                  `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                  `json:"@odata.type,omitempty"`
	OwnerAppIds          *[]string                `json:"ownerAppIds,omitempty"`
	Planner              *BusinessScenarioPlanner `json:"planner,omitempty"`
	UniqueName           *string                  `json:"uniqueName,omitempty"`
}
