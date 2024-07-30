package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalTaskSource struct {
	ContextScenarioId     *string                                      `json:"contextScenarioId,omitempty"`
	CreationSourceKind    *PlannerExternalTaskSourceCreationSourceKind `json:"creationSourceKind,omitempty"`
	DisplayLinkType       *PlannerExternalTaskSourceDisplayLinkType    `json:"displayLinkType,omitempty"`
	DisplayNameSegments   *[]string                                    `json:"displayNameSegments,omitempty"`
	ExternalContextId     *string                                      `json:"externalContextId,omitempty"`
	ExternalObjectId      *string                                      `json:"externalObjectId,omitempty"`
	ExternalObjectVersion *string                                      `json:"externalObjectVersion,omitempty"`
	ODataType             *string                                      `json:"@odata.type,omitempty"`
	TeamsPublicationInfo  *PlannerTeamsPublicationInfo                 `json:"teamsPublicationInfo,omitempty"`
	WebUrl                *string                                      `json:"webUrl,omitempty"`
}
