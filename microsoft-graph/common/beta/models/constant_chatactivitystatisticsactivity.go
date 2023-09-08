package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatActivityStatisticsActivity string

const (
	ChatActivityStatisticsActivityCall    ChatActivityStatisticsActivity = "Call"
	ChatActivityStatisticsActivityChat    ChatActivityStatisticsActivity = "Chat"
	ChatActivityStatisticsActivityEmail   ChatActivityStatisticsActivity = "Email"
	ChatActivityStatisticsActivityFocus   ChatActivityStatisticsActivity = "Focus"
	ChatActivityStatisticsActivityMeeting ChatActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForChatActivityStatisticsActivity() []string {
	return []string{
		string(ChatActivityStatisticsActivityCall),
		string(ChatActivityStatisticsActivityChat),
		string(ChatActivityStatisticsActivityEmail),
		string(ChatActivityStatisticsActivityFocus),
		string(ChatActivityStatisticsActivityMeeting),
	}
}

func parseChatActivityStatisticsActivity(input string) (*ChatActivityStatisticsActivity, error) {
	vals := map[string]ChatActivityStatisticsActivity{
		"call":    ChatActivityStatisticsActivityCall,
		"chat":    ChatActivityStatisticsActivityChat,
		"email":   ChatActivityStatisticsActivityEmail,
		"focus":   ChatActivityStatisticsActivityFocus,
		"meeting": ChatActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatActivityStatisticsActivity(input)
	return &out, nil
}
