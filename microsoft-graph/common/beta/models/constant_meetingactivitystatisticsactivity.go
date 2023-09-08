package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingActivityStatisticsActivity string

const (
	MeetingActivityStatisticsActivityCall    MeetingActivityStatisticsActivity = "Call"
	MeetingActivityStatisticsActivityChat    MeetingActivityStatisticsActivity = "Chat"
	MeetingActivityStatisticsActivityEmail   MeetingActivityStatisticsActivity = "Email"
	MeetingActivityStatisticsActivityFocus   MeetingActivityStatisticsActivity = "Focus"
	MeetingActivityStatisticsActivityMeeting MeetingActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForMeetingActivityStatisticsActivity() []string {
	return []string{
		string(MeetingActivityStatisticsActivityCall),
		string(MeetingActivityStatisticsActivityChat),
		string(MeetingActivityStatisticsActivityEmail),
		string(MeetingActivityStatisticsActivityFocus),
		string(MeetingActivityStatisticsActivityMeeting),
	}
}

func parseMeetingActivityStatisticsActivity(input string) (*MeetingActivityStatisticsActivity, error) {
	vals := map[string]MeetingActivityStatisticsActivity{
		"call":    MeetingActivityStatisticsActivityCall,
		"chat":    MeetingActivityStatisticsActivityChat,
		"email":   MeetingActivityStatisticsActivityEmail,
		"focus":   MeetingActivityStatisticsActivityFocus,
		"meeting": MeetingActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingActivityStatisticsActivity(input)
	return &out, nil
}
