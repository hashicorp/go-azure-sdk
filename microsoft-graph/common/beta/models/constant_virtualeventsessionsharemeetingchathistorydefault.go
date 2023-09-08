package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionShareMeetingChatHistoryDefault string

const (
	VirtualEventSessionShareMeetingChatHistoryDefaultall  VirtualEventSessionShareMeetingChatHistoryDefault = "All"
	VirtualEventSessionShareMeetingChatHistoryDefaultnone VirtualEventSessionShareMeetingChatHistoryDefault = "None"
)

func PossibleValuesForVirtualEventSessionShareMeetingChatHistoryDefault() []string {
	return []string{
		string(VirtualEventSessionShareMeetingChatHistoryDefaultall),
		string(VirtualEventSessionShareMeetingChatHistoryDefaultnone),
	}
}

func parseVirtualEventSessionShareMeetingChatHistoryDefault(input string) (*VirtualEventSessionShareMeetingChatHistoryDefault, error) {
	vals := map[string]VirtualEventSessionShareMeetingChatHistoryDefault{
		"all":  VirtualEventSessionShareMeetingChatHistoryDefaultall,
		"none": VirtualEventSessionShareMeetingChatHistoryDefaultnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionShareMeetingChatHistoryDefault(input)
	return &out, nil
}
