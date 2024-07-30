package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionAllowMeetingChat string

const (
	VirtualEventSessionAllowMeetingChat_Disabled VirtualEventSessionAllowMeetingChat = "disabled"
	VirtualEventSessionAllowMeetingChat_Enabled  VirtualEventSessionAllowMeetingChat = "enabled"
	VirtualEventSessionAllowMeetingChat_Limited  VirtualEventSessionAllowMeetingChat = "limited"
)

func PossibleValuesForVirtualEventSessionAllowMeetingChat() []string {
	return []string{
		string(VirtualEventSessionAllowMeetingChat_Disabled),
		string(VirtualEventSessionAllowMeetingChat_Enabled),
		string(VirtualEventSessionAllowMeetingChat_Limited),
	}
}

func (s *VirtualEventSessionAllowMeetingChat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionAllowMeetingChat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionAllowMeetingChat(input string) (*VirtualEventSessionAllowMeetingChat, error) {
	vals := map[string]VirtualEventSessionAllowMeetingChat{
		"disabled": VirtualEventSessionAllowMeetingChat_Disabled,
		"enabled":  VirtualEventSessionAllowMeetingChat_Enabled,
		"limited":  VirtualEventSessionAllowMeetingChat_Limited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAllowMeetingChat(input)
	return &out, nil
}
