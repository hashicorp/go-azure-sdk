package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBaseAllowMeetingChat string

const (
	OnlineMeetingBaseAllowMeetingChat_Disabled OnlineMeetingBaseAllowMeetingChat = "disabled"
	OnlineMeetingBaseAllowMeetingChat_Enabled  OnlineMeetingBaseAllowMeetingChat = "enabled"
	OnlineMeetingBaseAllowMeetingChat_Limited  OnlineMeetingBaseAllowMeetingChat = "limited"
)

func PossibleValuesForOnlineMeetingBaseAllowMeetingChat() []string {
	return []string{
		string(OnlineMeetingBaseAllowMeetingChat_Disabled),
		string(OnlineMeetingBaseAllowMeetingChat_Enabled),
		string(OnlineMeetingBaseAllowMeetingChat_Limited),
	}
}

func (s *OnlineMeetingBaseAllowMeetingChat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingBaseAllowMeetingChat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingBaseAllowMeetingChat(input string) (*OnlineMeetingBaseAllowMeetingChat, error) {
	vals := map[string]OnlineMeetingBaseAllowMeetingChat{
		"disabled": OnlineMeetingBaseAllowMeetingChat_Disabled,
		"enabled":  OnlineMeetingBaseAllowMeetingChat_Enabled,
		"limited":  OnlineMeetingBaseAllowMeetingChat_Limited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingBaseAllowMeetingChat(input)
	return &out, nil
}
