package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatActivityStatisticsActivity string

const (
	ChatActivityStatisticsActivity_Call    ChatActivityStatisticsActivity = "Call"
	ChatActivityStatisticsActivity_Chat    ChatActivityStatisticsActivity = "Chat"
	ChatActivityStatisticsActivity_Email   ChatActivityStatisticsActivity = "Email"
	ChatActivityStatisticsActivity_Focus   ChatActivityStatisticsActivity = "Focus"
	ChatActivityStatisticsActivity_Meeting ChatActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForChatActivityStatisticsActivity() []string {
	return []string{
		string(ChatActivityStatisticsActivity_Call),
		string(ChatActivityStatisticsActivity_Chat),
		string(ChatActivityStatisticsActivity_Email),
		string(ChatActivityStatisticsActivity_Focus),
		string(ChatActivityStatisticsActivity_Meeting),
	}
}

func (s *ChatActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatActivityStatisticsActivity(input string) (*ChatActivityStatisticsActivity, error) {
	vals := map[string]ChatActivityStatisticsActivity{
		"call":    ChatActivityStatisticsActivity_Call,
		"chat":    ChatActivityStatisticsActivity_Chat,
		"email":   ChatActivityStatisticsActivity_Email,
		"focus":   ChatActivityStatisticsActivity_Focus,
		"meeting": ChatActivityStatisticsActivity_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatActivityStatisticsActivity(input)
	return &out, nil
}
