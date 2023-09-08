package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettingsUserNewMessageRestriction string

const (
	ChannelModerationSettingsUserNewMessageRestrictioneveryone             ChannelModerationSettingsUserNewMessageRestriction = "Everyone"
	ChannelModerationSettingsUserNewMessageRestrictioneveryoneExceptGuests ChannelModerationSettingsUserNewMessageRestriction = "EveryoneExceptGuests"
	ChannelModerationSettingsUserNewMessageRestrictionmoderators           ChannelModerationSettingsUserNewMessageRestriction = "Moderators"
)

func PossibleValuesForChannelModerationSettingsUserNewMessageRestriction() []string {
	return []string{
		string(ChannelModerationSettingsUserNewMessageRestrictioneveryone),
		string(ChannelModerationSettingsUserNewMessageRestrictioneveryoneExceptGuests),
		string(ChannelModerationSettingsUserNewMessageRestrictionmoderators),
	}
}

func parseChannelModerationSettingsUserNewMessageRestriction(input string) (*ChannelModerationSettingsUserNewMessageRestriction, error) {
	vals := map[string]ChannelModerationSettingsUserNewMessageRestriction{
		"everyone":             ChannelModerationSettingsUserNewMessageRestrictioneveryone,
		"everyoneexceptguests": ChannelModerationSettingsUserNewMessageRestrictioneveryoneExceptGuests,
		"moderators":           ChannelModerationSettingsUserNewMessageRestrictionmoderators,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelModerationSettingsUserNewMessageRestriction(input)
	return &out, nil
}
