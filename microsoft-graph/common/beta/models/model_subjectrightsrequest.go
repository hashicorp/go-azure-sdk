package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequest struct {
	Approvers              *[]User                              `json:"approvers,omitempty"`
	AssignedTo             *Identity                            `json:"assignedTo,omitempty"`
	ClosedDateTime         *string                              `json:"closedDateTime,omitempty"`
	Collaborators          *[]User                              `json:"collaborators,omitempty"`
	ContentQuery           *string                              `json:"contentQuery,omitempty"`
	CreatedBy              *IdentitySet                         `json:"createdBy,omitempty"`
	CreatedDateTime        *string                              `json:"createdDateTime,omitempty"`
	DataSubject            *DataSubject                         `json:"dataSubject,omitempty"`
	DataSubjectType        *SubjectRightsRequestDataSubjectType `json:"dataSubjectType,omitempty"`
	Description            *string                              `json:"description,omitempty"`
	DisplayName            *string                              `json:"displayName,omitempty"`
	ExternalId             *string                              `json:"externalId,omitempty"`
	History                *[]SubjectRightsRequestHistory       `json:"history,omitempty"`
	Id                     *string                              `json:"id,omitempty"`
	IncludeAllVersions     *bool                                `json:"includeAllVersions,omitempty"`
	IncludeAuthoredContent *bool                                `json:"includeAuthoredContent,omitempty"`
	Insight                *SubjectRightsRequestDetail          `json:"insight,omitempty"`
	InternalDueDateTime    *string                              `json:"internalDueDateTime,omitempty"`
	LastModifiedBy         *IdentitySet                         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime   *string                              `json:"lastModifiedDateTime,omitempty"`
	Mailboxlocations       *SubjectRightsRequestMailboxLocation `json:"mailboxlocations,omitempty"`
	Notes                  *[]AuthoredNote                      `json:"notes,omitempty"`
	ODataType              *string                              `json:"@odata.type,omitempty"`
	PauseAfterEstimate     *bool                                `json:"pauseAfterEstimate,omitempty"`
	Regulations            *[]string                            `json:"regulations,omitempty"`
	Sitelocations          *SubjectRightsRequestSiteLocation    `json:"sitelocations,omitempty"`
	Stages                 *[]SubjectRightsRequestStageDetail   `json:"stages,omitempty"`
	Status                 *SubjectRightsRequestStatus          `json:"status,omitempty"`
	Team                   *Team                                `json:"team,omitempty"`
	Type                   *SubjectRightsRequestType            `json:"type,omitempty"`
}
