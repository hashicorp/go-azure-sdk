package joinedteam

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateJoinedTeamCloneRequest struct {
	Classification nullable.Type[string]      `json:"classification,omitempty"`
	Description    nullable.Type[string]      `json:"description,omitempty"`
	DisplayName    nullable.Type[string]      `json:"displayName,omitempty"`
	MailNickname   nullable.Type[string]      `json:"mailNickname,omitempty"`
	PartsToClone   *stable.ClonableTeamParts  `json:"partsToClone,omitempty"`
	Visibility     *stable.TeamVisibilityType `json:"visibility,omitempty"`
}
