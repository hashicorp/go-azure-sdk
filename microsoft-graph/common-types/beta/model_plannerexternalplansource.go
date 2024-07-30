package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalPlanSource struct {
	ContextScenarioId  *string                                      `json:"contextScenarioId,omitempty"`
	CreationSourceKind *PlannerExternalPlanSourceCreationSourceKind `json:"creationSourceKind,omitempty"`
	ExternalContextId  *string                                      `json:"externalContextId,omitempty"`
	ExternalObjectId   *string                                      `json:"externalObjectId,omitempty"`
	ODataType          *string                                      `json:"@odata.type,omitempty"`
}
