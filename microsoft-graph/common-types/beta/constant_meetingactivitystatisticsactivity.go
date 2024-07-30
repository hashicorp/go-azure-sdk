package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingActivityStatisticsActivity string

const (
	MeetingActivityStatisticsActivity_Call    MeetingActivityStatisticsActivity = "Call"
	MeetingActivityStatisticsActivity_Chat    MeetingActivityStatisticsActivity = "Chat"
	MeetingActivityStatisticsActivity_Email   MeetingActivityStatisticsActivity = "Email"
	MeetingActivityStatisticsActivity_Focus   MeetingActivityStatisticsActivity = "Focus"
	MeetingActivityStatisticsActivity_Meeting MeetingActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForMeetingActivityStatisticsActivity() []string {
	return []string{
		string(MeetingActivityStatisticsActivity_Call),
		string(MeetingActivityStatisticsActivity_Chat),
		string(MeetingActivityStatisticsActivity_Email),
		string(MeetingActivityStatisticsActivity_Focus),
		string(MeetingActivityStatisticsActivity_Meeting),
	}
}

func (s *MeetingActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingActivityStatisticsActivity(input string) (*MeetingActivityStatisticsActivity, error) {
	vals := map[string]MeetingActivityStatisticsActivity{
		"call":    MeetingActivityStatisticsActivity_Call,
		"chat":    MeetingActivityStatisticsActivity_Chat,
		"email":   MeetingActivityStatisticsActivity_Email,
		"focus":   MeetingActivityStatisticsActivity_Focus,
		"meeting": MeetingActivityStatisticsActivity_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingActivityStatisticsActivity(input)
	return &out, nil
}
