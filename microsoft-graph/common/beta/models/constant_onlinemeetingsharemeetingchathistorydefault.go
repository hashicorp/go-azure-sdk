package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingShareMeetingChatHistoryDefault string

const (
	OnlineMeetingShareMeetingChatHistoryDefaultall  OnlineMeetingShareMeetingChatHistoryDefault = "All"
	OnlineMeetingShareMeetingChatHistoryDefaultnone OnlineMeetingShareMeetingChatHistoryDefault = "None"
)

func PossibleValuesForOnlineMeetingShareMeetingChatHistoryDefault() []string {
	return []string{
		string(OnlineMeetingShareMeetingChatHistoryDefaultall),
		string(OnlineMeetingShareMeetingChatHistoryDefaultnone),
	}
}

func parseOnlineMeetingShareMeetingChatHistoryDefault(input string) (*OnlineMeetingShareMeetingChatHistoryDefault, error) {
	vals := map[string]OnlineMeetingShareMeetingChatHistoryDefault{
		"all":  OnlineMeetingShareMeetingChatHistoryDefaultall,
		"none": OnlineMeetingShareMeetingChatHistoryDefaultnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingShareMeetingChatHistoryDefault(input)
	return &out, nil
}
