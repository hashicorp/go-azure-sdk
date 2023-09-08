package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityStatisticsActivity string

const (
	ActivityStatisticsActivityCall    ActivityStatisticsActivity = "Call"
	ActivityStatisticsActivityChat    ActivityStatisticsActivity = "Chat"
	ActivityStatisticsActivityEmail   ActivityStatisticsActivity = "Email"
	ActivityStatisticsActivityFocus   ActivityStatisticsActivity = "Focus"
	ActivityStatisticsActivityMeeting ActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForActivityStatisticsActivity() []string {
	return []string{
		string(ActivityStatisticsActivityCall),
		string(ActivityStatisticsActivityChat),
		string(ActivityStatisticsActivityEmail),
		string(ActivityStatisticsActivityFocus),
		string(ActivityStatisticsActivityMeeting),
	}
}

func parseActivityStatisticsActivity(input string) (*ActivityStatisticsActivity, error) {
	vals := map[string]ActivityStatisticsActivity{
		"call":    ActivityStatisticsActivityCall,
		"chat":    ActivityStatisticsActivityChat,
		"email":   ActivityStatisticsActivityEmail,
		"focus":   ActivityStatisticsActivityFocus,
		"meeting": ActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityStatisticsActivity(input)
	return &out, nil
}
