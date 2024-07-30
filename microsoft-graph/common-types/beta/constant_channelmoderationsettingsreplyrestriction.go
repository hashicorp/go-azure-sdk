package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettingsReplyRestriction string

const (
	ChannelModerationSettingsReplyRestriction_AuthorAndModerators ChannelModerationSettingsReplyRestriction = "authorAndModerators"
	ChannelModerationSettingsReplyRestriction_Everyone            ChannelModerationSettingsReplyRestriction = "everyone"
)

func PossibleValuesForChannelModerationSettingsReplyRestriction() []string {
	return []string{
		string(ChannelModerationSettingsReplyRestriction_AuthorAndModerators),
		string(ChannelModerationSettingsReplyRestriction_Everyone),
	}
}

func (s *ChannelModerationSettingsReplyRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelModerationSettingsReplyRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelModerationSettingsReplyRestriction(input string) (*ChannelModerationSettingsReplyRestriction, error) {
	vals := map[string]ChannelModerationSettingsReplyRestriction{
		"authorandmoderators": ChannelModerationSettingsReplyRestriction_AuthorAndModerators,
		"everyone":            ChannelModerationSettingsReplyRestriction_Everyone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelModerationSettingsReplyRestriction(input)
	return &out, nil
}
