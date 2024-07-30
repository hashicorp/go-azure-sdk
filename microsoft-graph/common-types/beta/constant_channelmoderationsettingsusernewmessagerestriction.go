package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettingsUserNewMessageRestriction string

const (
	ChannelModerationSettingsUserNewMessageRestriction_Everyone             ChannelModerationSettingsUserNewMessageRestriction = "everyone"
	ChannelModerationSettingsUserNewMessageRestriction_EveryoneExceptGuests ChannelModerationSettingsUserNewMessageRestriction = "everyoneExceptGuests"
	ChannelModerationSettingsUserNewMessageRestriction_Moderators           ChannelModerationSettingsUserNewMessageRestriction = "moderators"
)

func PossibleValuesForChannelModerationSettingsUserNewMessageRestriction() []string {
	return []string{
		string(ChannelModerationSettingsUserNewMessageRestriction_Everyone),
		string(ChannelModerationSettingsUserNewMessageRestriction_EveryoneExceptGuests),
		string(ChannelModerationSettingsUserNewMessageRestriction_Moderators),
	}
}

func (s *ChannelModerationSettingsUserNewMessageRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelModerationSettingsUserNewMessageRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelModerationSettingsUserNewMessageRestriction(input string) (*ChannelModerationSettingsUserNewMessageRestriction, error) {
	vals := map[string]ChannelModerationSettingsUserNewMessageRestriction{
		"everyone":             ChannelModerationSettingsUserNewMessageRestriction_Everyone,
		"everyoneexceptguests": ChannelModerationSettingsUserNewMessageRestriction_EveryoneExceptGuests,
		"moderators":           ChannelModerationSettingsUserNewMessageRestriction_Moderators,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelModerationSettingsUserNewMessageRestriction(input)
	return &out, nil
}
