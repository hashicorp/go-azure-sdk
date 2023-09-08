package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettingsReplyRestriction string

const (
	ChannelModerationSettingsReplyRestrictionauthorAndModerators ChannelModerationSettingsReplyRestriction = "AuthorAndModerators"
	ChannelModerationSettingsReplyRestrictioneveryone            ChannelModerationSettingsReplyRestriction = "Everyone"
)

func PossibleValuesForChannelModerationSettingsReplyRestriction() []string {
	return []string{
		string(ChannelModerationSettingsReplyRestrictionauthorAndModerators),
		string(ChannelModerationSettingsReplyRestrictioneveryone),
	}
}

func parseChannelModerationSettingsReplyRestriction(input string) (*ChannelModerationSettingsReplyRestriction, error) {
	vals := map[string]ChannelModerationSettingsReplyRestriction{
		"authorandmoderators": ChannelModerationSettingsReplyRestrictionauthorAndModerators,
		"everyone":            ChannelModerationSettingsReplyRestrictioneveryone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelModerationSettingsReplyRestriction(input)
	return &out, nil
}
