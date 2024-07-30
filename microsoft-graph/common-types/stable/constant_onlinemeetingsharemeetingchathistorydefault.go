package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingShareMeetingChatHistoryDefault string

const (
	OnlineMeetingShareMeetingChatHistoryDefault_All  OnlineMeetingShareMeetingChatHistoryDefault = "all"
	OnlineMeetingShareMeetingChatHistoryDefault_None OnlineMeetingShareMeetingChatHistoryDefault = "none"
)

func PossibleValuesForOnlineMeetingShareMeetingChatHistoryDefault() []string {
	return []string{
		string(OnlineMeetingShareMeetingChatHistoryDefault_All),
		string(OnlineMeetingShareMeetingChatHistoryDefault_None),
	}
}

func (s *OnlineMeetingShareMeetingChatHistoryDefault) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingShareMeetingChatHistoryDefault(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingShareMeetingChatHistoryDefault(input string) (*OnlineMeetingShareMeetingChatHistoryDefault, error) {
	vals := map[string]OnlineMeetingShareMeetingChatHistoryDefault{
		"all":  OnlineMeetingShareMeetingChatHistoryDefault_All,
		"none": OnlineMeetingShareMeetingChatHistoryDefault_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingShareMeetingChatHistoryDefault(input)
	return &out, nil
}
