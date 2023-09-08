package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditions struct {
	AcceptanceStatement  *string                               `json:"acceptanceStatement,omitempty"`
	AcceptanceStatuses   *[]TermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`
	Assignments          *[]TermsAndConditionsAssignment       `json:"assignments,omitempty"`
	BodyText             *string                               `json:"bodyText,omitempty"`
	CreatedDateTime      *string                               `json:"createdDateTime,omitempty"`
	Description          *string                               `json:"description,omitempty"`
	DisplayName          *string                               `json:"displayName,omitempty"`
	GroupAssignments     *[]TermsAndConditionsGroupAssignment  `json:"groupAssignments,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	LastModifiedDateTime *string                               `json:"lastModifiedDateTime,omitempty"`
	ModifiedDateTime     *string                               `json:"modifiedDateTime,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	RoleScopeTagIds      *[]string                             `json:"roleScopeTagIds,omitempty"`
	Title                *string                               `json:"title,omitempty"`
	Version              *int64                                `json:"version,omitempty"`
}
