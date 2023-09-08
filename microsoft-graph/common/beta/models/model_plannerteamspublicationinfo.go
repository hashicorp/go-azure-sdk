package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTeamsPublicationInfo struct {
	CreationSourceKind   *PlannerTeamsPublicationInfoCreationSourceKind `json:"creationSourceKind,omitempty"`
	LastModifiedDateTime *string                                        `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                        `json:"@odata.type,omitempty"`
	PublicationId        *string                                        `json:"publicationId,omitempty"`
	PublishedToPlanId    *string                                        `json:"publishedToPlanId,omitempty"`
	PublishingTeamId     *string                                        `json:"publishingTeamId,omitempty"`
	PublishingTeamName   *string                                        `json:"publishingTeamName,omitempty"`
	TeamsPublicationInfo *PlannerTeamsPublicationInfo                   `json:"teamsPublicationInfo,omitempty"`
}
