package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConversationIdentityConversationIdentityType string

const (
	TeamworkConversationIdentityConversationIdentityType_Channel TeamworkConversationIdentityConversationIdentityType = "channel"
	TeamworkConversationIdentityConversationIdentityType_Chat    TeamworkConversationIdentityConversationIdentityType = "chat"
	TeamworkConversationIdentityConversationIdentityType_Team    TeamworkConversationIdentityConversationIdentityType = "team"
)

func PossibleValuesForTeamworkConversationIdentityConversationIdentityType() []string {
	return []string{
		string(TeamworkConversationIdentityConversationIdentityType_Channel),
		string(TeamworkConversationIdentityConversationIdentityType_Chat),
		string(TeamworkConversationIdentityConversationIdentityType_Team),
	}
}

func (s *TeamworkConversationIdentityConversationIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkConversationIdentityConversationIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkConversationIdentityConversationIdentityType(input string) (*TeamworkConversationIdentityConversationIdentityType, error) {
	vals := map[string]TeamworkConversationIdentityConversationIdentityType{
		"channel": TeamworkConversationIdentityConversationIdentityType_Channel,
		"chat":    TeamworkConversationIdentityConversationIdentityType_Chat,
		"team":    TeamworkConversationIdentityConversationIdentityType_Team,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConversationIdentityConversationIdentityType(input)
	return &out, nil
}
