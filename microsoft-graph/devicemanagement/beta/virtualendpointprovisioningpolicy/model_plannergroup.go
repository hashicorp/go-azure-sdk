package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerGroup struct {
	Id        *string        `json:"id,omitempty"`
	ODataType *string        `json:"@odata.type,omitempty"`
	Plans     *[]PlannerPlan `json:"plans,omitempty"`
}
