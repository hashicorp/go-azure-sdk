package mejoinedteam

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeJoinedTeamByIdCloneRequest struct {
	Classification *string                                         `json:"classification,omitempty"`
	Description    *string                                         `json:"description,omitempty"`
	DisplayName    *string                                         `json:"displayName,omitempty"`
	MailNickname   *string                                         `json:"mailNickname,omitempty"`
	PartsToClone   *CreateMeJoinedTeamByIdCloneRequestPartsToClone `json:"partsToClone,omitempty"`
	Visibility     *CreateMeJoinedTeamByIdCloneRequestVisibility   `json:"visibility,omitempty"`
}
