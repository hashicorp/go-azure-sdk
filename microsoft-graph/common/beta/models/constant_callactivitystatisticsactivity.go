package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallActivityStatisticsActivity string

const (
	CallActivityStatisticsActivityCall    CallActivityStatisticsActivity = "Call"
	CallActivityStatisticsActivityChat    CallActivityStatisticsActivity = "Chat"
	CallActivityStatisticsActivityEmail   CallActivityStatisticsActivity = "Email"
	CallActivityStatisticsActivityFocus   CallActivityStatisticsActivity = "Focus"
	CallActivityStatisticsActivityMeeting CallActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForCallActivityStatisticsActivity() []string {
	return []string{
		string(CallActivityStatisticsActivityCall),
		string(CallActivityStatisticsActivityChat),
		string(CallActivityStatisticsActivityEmail),
		string(CallActivityStatisticsActivityFocus),
		string(CallActivityStatisticsActivityMeeting),
	}
}

func parseCallActivityStatisticsActivity(input string) (*CallActivityStatisticsActivity, error) {
	vals := map[string]CallActivityStatisticsActivity{
		"call":    CallActivityStatisticsActivityCall,
		"chat":    CallActivityStatisticsActivityChat,
		"email":   CallActivityStatisticsActivityEmail,
		"focus":   CallActivityStatisticsActivityFocus,
		"meeting": CallActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallActivityStatisticsActivity(input)
	return &out, nil
}
