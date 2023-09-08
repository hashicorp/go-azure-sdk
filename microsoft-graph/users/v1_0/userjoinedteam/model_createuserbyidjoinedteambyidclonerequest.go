package userjoinedteam

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdJoinedTeamByIdCloneRequest struct {
	Classification *string                                               `json:"classification,omitempty"`
	Description    *string                                               `json:"description,omitempty"`
	DisplayName    *string                                               `json:"displayName,omitempty"`
	MailNickname   *string                                               `json:"mailNickname,omitempty"`
	PartsToClone   *CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone `json:"partsToClone,omitempty"`
	Visibility     *CreateUserByIdJoinedTeamByIdCloneRequestVisibility   `json:"visibility,omitempty"`
}
