package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanCreation struct {
	CreationSourceKind *PlannerPlanCreationCreationSourceKind `json:"creationSourceKind,omitempty"`
	ODataType          *string                                `json:"@odata.type,omitempty"`
}
