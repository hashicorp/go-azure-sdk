package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskCreation struct {
	CreationSourceKind   *PlannerTaskCreationCreationSourceKind `json:"creationSourceKind,omitempty"`
	ODataType            *string                                `json:"@odata.type,omitempty"`
	TeamsPublicationInfo *PlannerTeamsPublicationInfo           `json:"teamsPublicationInfo,omitempty"`
}
