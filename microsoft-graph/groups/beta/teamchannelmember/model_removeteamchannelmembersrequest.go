package teamchannelmember

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveTeamChannelMembersRequest struct {
	Values *[]beta.ConversationMember `json:"values,omitempty"`
}
