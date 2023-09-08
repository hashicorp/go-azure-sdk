package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailActivityStatisticsActivity string

const (
	EmailActivityStatisticsActivityCall    EmailActivityStatisticsActivity = "Call"
	EmailActivityStatisticsActivityChat    EmailActivityStatisticsActivity = "Chat"
	EmailActivityStatisticsActivityEmail   EmailActivityStatisticsActivity = "Email"
	EmailActivityStatisticsActivityFocus   EmailActivityStatisticsActivity = "Focus"
	EmailActivityStatisticsActivityMeeting EmailActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForEmailActivityStatisticsActivity() []string {
	return []string{
		string(EmailActivityStatisticsActivityCall),
		string(EmailActivityStatisticsActivityChat),
		string(EmailActivityStatisticsActivityEmail),
		string(EmailActivityStatisticsActivityFocus),
		string(EmailActivityStatisticsActivityMeeting),
	}
}

func parseEmailActivityStatisticsActivity(input string) (*EmailActivityStatisticsActivity, error) {
	vals := map[string]EmailActivityStatisticsActivity{
		"call":    EmailActivityStatisticsActivityCall,
		"chat":    EmailActivityStatisticsActivityChat,
		"email":   EmailActivityStatisticsActivityEmail,
		"focus":   EmailActivityStatisticsActivityFocus,
		"meeting": EmailActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailActivityStatisticsActivity(input)
	return &out, nil
}
