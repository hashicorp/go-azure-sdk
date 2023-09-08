package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelMembershipType string

const (
	ChannelMembershipTypeprivate  ChannelMembershipType = "Private"
	ChannelMembershipTypeshared   ChannelMembershipType = "Shared"
	ChannelMembershipTypestandard ChannelMembershipType = "Standard"
)

func PossibleValuesForChannelMembershipType() []string {
	return []string{
		string(ChannelMembershipTypeprivate),
		string(ChannelMembershipTypeshared),
		string(ChannelMembershipTypestandard),
	}
}

func parseChannelMembershipType(input string) (*ChannelMembershipType, error) {
	vals := map[string]ChannelMembershipType{
		"private":  ChannelMembershipTypeprivate,
		"shared":   ChannelMembershipTypeshared,
		"standard": ChannelMembershipTypestandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelMembershipType(input)
	return &out, nil
}
