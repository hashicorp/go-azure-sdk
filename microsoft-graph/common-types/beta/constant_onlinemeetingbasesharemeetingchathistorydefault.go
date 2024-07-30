package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBaseShareMeetingChatHistoryDefault string

const (
	OnlineMeetingBaseShareMeetingChatHistoryDefault_All  OnlineMeetingBaseShareMeetingChatHistoryDefault = "all"
	OnlineMeetingBaseShareMeetingChatHistoryDefault_None OnlineMeetingBaseShareMeetingChatHistoryDefault = "none"
)

func PossibleValuesForOnlineMeetingBaseShareMeetingChatHistoryDefault() []string {
	return []string{
		string(OnlineMeetingBaseShareMeetingChatHistoryDefault_All),
		string(OnlineMeetingBaseShareMeetingChatHistoryDefault_None),
	}
}

func (s *OnlineMeetingBaseShareMeetingChatHistoryDefault) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingBaseShareMeetingChatHistoryDefault(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingBaseShareMeetingChatHistoryDefault(input string) (*OnlineMeetingBaseShareMeetingChatHistoryDefault, error) {
	vals := map[string]OnlineMeetingBaseShareMeetingChatHistoryDefault{
		"all":  OnlineMeetingBaseShareMeetingChatHistoryDefault_All,
		"none": OnlineMeetingBaseShareMeetingChatHistoryDefault_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingBaseShareMeetingChatHistoryDefault(input)
	return &out, nil
}
