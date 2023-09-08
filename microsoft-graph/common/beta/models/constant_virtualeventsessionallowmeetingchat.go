package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionAllowMeetingChat string

const (
	VirtualEventSessionAllowMeetingChatdisabled VirtualEventSessionAllowMeetingChat = "Disabled"
	VirtualEventSessionAllowMeetingChatenabled  VirtualEventSessionAllowMeetingChat = "Enabled"
	VirtualEventSessionAllowMeetingChatlimited  VirtualEventSessionAllowMeetingChat = "Limited"
)

func PossibleValuesForVirtualEventSessionAllowMeetingChat() []string {
	return []string{
		string(VirtualEventSessionAllowMeetingChatdisabled),
		string(VirtualEventSessionAllowMeetingChatenabled),
		string(VirtualEventSessionAllowMeetingChatlimited),
	}
}

func parseVirtualEventSessionAllowMeetingChat(input string) (*VirtualEventSessionAllowMeetingChat, error) {
	vals := map[string]VirtualEventSessionAllowMeetingChat{
		"disabled": VirtualEventSessionAllowMeetingChatdisabled,
		"enabled":  VirtualEventSessionAllowMeetingChatenabled,
		"limited":  VirtualEventSessionAllowMeetingChatlimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAllowMeetingChat(input)
	return &out, nil
}
