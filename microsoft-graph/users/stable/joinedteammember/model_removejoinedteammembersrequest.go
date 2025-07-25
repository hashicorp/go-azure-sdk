package joinedteammember

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveJoinedTeamMembersRequest struct {
	Values *[]stable.ConversationMember `json:"values,omitempty"`
}
