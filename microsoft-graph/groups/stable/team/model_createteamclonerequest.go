package team

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamCloneRequest struct {
	Classification *string                             `json:"classification,omitempty"`
	Description    *string                             `json:"description,omitempty"`
	DisplayName    *string                             `json:"displayName,omitempty"`
	MailNickname   *string                             `json:"mailNickname,omitempty"`
	PartsToClone   *CreateTeamCloneRequestPartsToClone `json:"partsToClone,omitempty"`
	Visibility     *CreateTeamCloneRequestVisibility   `json:"visibility,omitempty"`
}
