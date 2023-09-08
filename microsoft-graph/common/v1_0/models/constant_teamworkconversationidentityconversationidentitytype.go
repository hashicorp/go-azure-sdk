package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConversationIdentityConversationIdentityType string

const (
	TeamworkConversationIdentityConversationIdentityTypechannel TeamworkConversationIdentityConversationIdentityType = "Channel"
	TeamworkConversationIdentityConversationIdentityTypechat    TeamworkConversationIdentityConversationIdentityType = "Chat"
	TeamworkConversationIdentityConversationIdentityTypeteam    TeamworkConversationIdentityConversationIdentityType = "Team"
)

func PossibleValuesForTeamworkConversationIdentityConversationIdentityType() []string {
	return []string{
		string(TeamworkConversationIdentityConversationIdentityTypechannel),
		string(TeamworkConversationIdentityConversationIdentityTypechat),
		string(TeamworkConversationIdentityConversationIdentityTypeteam),
	}
}

func parseTeamworkConversationIdentityConversationIdentityType(input string) (*TeamworkConversationIdentityConversationIdentityType, error) {
	vals := map[string]TeamworkConversationIdentityConversationIdentityType{
		"channel": TeamworkConversationIdentityConversationIdentityTypechannel,
		"chat":    TeamworkConversationIdentityConversationIdentityTypechat,
		"team":    TeamworkConversationIdentityConversationIdentityTypeteam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConversationIdentityConversationIdentityType(input)
	return &out, nil
}
