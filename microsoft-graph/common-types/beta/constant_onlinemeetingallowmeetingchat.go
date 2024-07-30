package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAllowMeetingChat string

const (
	OnlineMeetingAllowMeetingChat_Disabled OnlineMeetingAllowMeetingChat = "disabled"
	OnlineMeetingAllowMeetingChat_Enabled  OnlineMeetingAllowMeetingChat = "enabled"
	OnlineMeetingAllowMeetingChat_Limited  OnlineMeetingAllowMeetingChat = "limited"
)

func PossibleValuesForOnlineMeetingAllowMeetingChat() []string {
	return []string{
		string(OnlineMeetingAllowMeetingChat_Disabled),
		string(OnlineMeetingAllowMeetingChat_Enabled),
		string(OnlineMeetingAllowMeetingChat_Limited),
	}
}

func (s *OnlineMeetingAllowMeetingChat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingAllowMeetingChat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingAllowMeetingChat(input string) (*OnlineMeetingAllowMeetingChat, error) {
	vals := map[string]OnlineMeetingAllowMeetingChat{
		"disabled": OnlineMeetingAllowMeetingChat_Disabled,
		"enabled":  OnlineMeetingAllowMeetingChat_Enabled,
		"limited":  OnlineMeetingAllowMeetingChat_Limited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAllowMeetingChat(input)
	return &out, nil
}
