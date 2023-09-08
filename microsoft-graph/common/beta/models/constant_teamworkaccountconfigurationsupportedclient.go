package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkAccountConfigurationSupportedClient string

const (
	TeamworkAccountConfigurationSupportedClientskypeDefaultAndTeams TeamworkAccountConfigurationSupportedClient = "SkypeDefaultAndTeams"
	TeamworkAccountConfigurationSupportedClientskypeOnly            TeamworkAccountConfigurationSupportedClient = "SkypeOnly"
	TeamworkAccountConfigurationSupportedClientteamsDefaultAndSkype TeamworkAccountConfigurationSupportedClient = "TeamsDefaultAndSkype"
	TeamworkAccountConfigurationSupportedClientteamsOnly            TeamworkAccountConfigurationSupportedClient = "TeamsOnly"
	TeamworkAccountConfigurationSupportedClientunknown              TeamworkAccountConfigurationSupportedClient = "Unknown"
)

func PossibleValuesForTeamworkAccountConfigurationSupportedClient() []string {
	return []string{
		string(TeamworkAccountConfigurationSupportedClientskypeDefaultAndTeams),
		string(TeamworkAccountConfigurationSupportedClientskypeOnly),
		string(TeamworkAccountConfigurationSupportedClientteamsDefaultAndSkype),
		string(TeamworkAccountConfigurationSupportedClientteamsOnly),
		string(TeamworkAccountConfigurationSupportedClientunknown),
	}
}

func parseTeamworkAccountConfigurationSupportedClient(input string) (*TeamworkAccountConfigurationSupportedClient, error) {
	vals := map[string]TeamworkAccountConfigurationSupportedClient{
		"skypedefaultandteams": TeamworkAccountConfigurationSupportedClientskypeDefaultAndTeams,
		"skypeonly":            TeamworkAccountConfigurationSupportedClientskypeOnly,
		"teamsdefaultandskype": TeamworkAccountConfigurationSupportedClientteamsDefaultAndSkype,
		"teamsonly":            TeamworkAccountConfigurationSupportedClientteamsOnly,
		"unknown":              TeamworkAccountConfigurationSupportedClientunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkAccountConfigurationSupportedClient(input)
	return &out, nil
}
