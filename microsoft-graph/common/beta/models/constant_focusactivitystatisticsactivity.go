package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FocusActivityStatisticsActivity string

const (
	FocusActivityStatisticsActivityCall    FocusActivityStatisticsActivity = "Call"
	FocusActivityStatisticsActivityChat    FocusActivityStatisticsActivity = "Chat"
	FocusActivityStatisticsActivityEmail   FocusActivityStatisticsActivity = "Email"
	FocusActivityStatisticsActivityFocus   FocusActivityStatisticsActivity = "Focus"
	FocusActivityStatisticsActivityMeeting FocusActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForFocusActivityStatisticsActivity() []string {
	return []string{
		string(FocusActivityStatisticsActivityCall),
		string(FocusActivityStatisticsActivityChat),
		string(FocusActivityStatisticsActivityEmail),
		string(FocusActivityStatisticsActivityFocus),
		string(FocusActivityStatisticsActivityMeeting),
	}
}

func parseFocusActivityStatisticsActivity(input string) (*FocusActivityStatisticsActivity, error) {
	vals := map[string]FocusActivityStatisticsActivity{
		"call":    FocusActivityStatisticsActivityCall,
		"chat":    FocusActivityStatisticsActivityChat,
		"email":   FocusActivityStatisticsActivityEmail,
		"focus":   FocusActivityStatisticsActivityFocus,
		"meeting": FocusActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FocusActivityStatisticsActivity(input)
	return &out, nil
}
