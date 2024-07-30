package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContext struct {
	AssociationType     *string   `json:"associationType,omitempty"`
	CreatedDateTime     *string   `json:"createdDateTime,omitempty"`
	DisplayNameSegments *[]string `json:"displayNameSegments,omitempty"`
	IsCreationContext   *bool     `json:"isCreationContext,omitempty"`
	ODataType           *string   `json:"@odata.type,omitempty"`
	OwnerAppId          *string   `json:"ownerAppId,omitempty"`
}
