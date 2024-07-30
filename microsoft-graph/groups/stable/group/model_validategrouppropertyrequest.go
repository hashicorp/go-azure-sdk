package group

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateGroupPropertyRequest struct {
	DisplayName      *string `json:"displayName,omitempty"`
	MailNickname     *string `json:"mailNickname,omitempty"`
	OnBehalfOfUserId *string `json:"onBehalfOfUserId,omitempty"`
}
