package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAllowMeetingChat string

const (
	OnlineMeetingAllowMeetingChatdisabled OnlineMeetingAllowMeetingChat = "Disabled"
	OnlineMeetingAllowMeetingChatenabled  OnlineMeetingAllowMeetingChat = "Enabled"
	OnlineMeetingAllowMeetingChatlimited  OnlineMeetingAllowMeetingChat = "Limited"
)

func PossibleValuesForOnlineMeetingAllowMeetingChat() []string {
	return []string{
		string(OnlineMeetingAllowMeetingChatdisabled),
		string(OnlineMeetingAllowMeetingChatenabled),
		string(OnlineMeetingAllowMeetingChatlimited),
	}
}

func parseOnlineMeetingAllowMeetingChat(input string) (*OnlineMeetingAllowMeetingChat, error) {
	vals := map[string]OnlineMeetingAllowMeetingChat{
		"disabled": OnlineMeetingAllowMeetingChatdisabled,
		"enabled":  OnlineMeetingAllowMeetingChatenabled,
		"limited":  OnlineMeetingAllowMeetingChatlimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAllowMeetingChat(input)
	return &out, nil
}
