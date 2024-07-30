package joinedteamchannelmember

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddUserJoinedTeamChannelMemberRequest struct {
	Values *[]ConversationMember `json:"values,omitempty"`
}
