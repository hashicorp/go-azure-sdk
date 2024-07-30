package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionShareMeetingChatHistoryDefault string

const (
	VirtualEventSessionShareMeetingChatHistoryDefault_All  VirtualEventSessionShareMeetingChatHistoryDefault = "all"
	VirtualEventSessionShareMeetingChatHistoryDefault_None VirtualEventSessionShareMeetingChatHistoryDefault = "none"
)

func PossibleValuesForVirtualEventSessionShareMeetingChatHistoryDefault() []string {
	return []string{
		string(VirtualEventSessionShareMeetingChatHistoryDefault_All),
		string(VirtualEventSessionShareMeetingChatHistoryDefault_None),
	}
}

func (s *VirtualEventSessionShareMeetingChatHistoryDefault) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionShareMeetingChatHistoryDefault(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionShareMeetingChatHistoryDefault(input string) (*VirtualEventSessionShareMeetingChatHistoryDefault, error) {
	vals := map[string]VirtualEventSessionShareMeetingChatHistoryDefault{
		"all":  VirtualEventSessionShareMeetingChatHistoryDefault_All,
		"none": VirtualEventSessionShareMeetingChatHistoryDefault_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionShareMeetingChatHistoryDefault(input)
	return &out, nil
}
