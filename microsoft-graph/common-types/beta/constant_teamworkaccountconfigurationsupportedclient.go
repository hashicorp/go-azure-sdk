package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkAccountConfigurationSupportedClient string

const (
	TeamworkAccountConfigurationSupportedClient_SkypeDefaultAndTeams TeamworkAccountConfigurationSupportedClient = "skypeDefaultAndTeams"
	TeamworkAccountConfigurationSupportedClient_SkypeOnly            TeamworkAccountConfigurationSupportedClient = "skypeOnly"
	TeamworkAccountConfigurationSupportedClient_TeamsDefaultAndSkype TeamworkAccountConfigurationSupportedClient = "teamsDefaultAndSkype"
	TeamworkAccountConfigurationSupportedClient_TeamsOnly            TeamworkAccountConfigurationSupportedClient = "teamsOnly"
	TeamworkAccountConfigurationSupportedClient_Unknown              TeamworkAccountConfigurationSupportedClient = "unknown"
)

func PossibleValuesForTeamworkAccountConfigurationSupportedClient() []string {
	return []string{
		string(TeamworkAccountConfigurationSupportedClient_SkypeDefaultAndTeams),
		string(TeamworkAccountConfigurationSupportedClient_SkypeOnly),
		string(TeamworkAccountConfigurationSupportedClient_TeamsDefaultAndSkype),
		string(TeamworkAccountConfigurationSupportedClient_TeamsOnly),
		string(TeamworkAccountConfigurationSupportedClient_Unknown),
	}
}

func (s *TeamworkAccountConfigurationSupportedClient) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkAccountConfigurationSupportedClient(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkAccountConfigurationSupportedClient(input string) (*TeamworkAccountConfigurationSupportedClient, error) {
	vals := map[string]TeamworkAccountConfigurationSupportedClient{
		"skypedefaultandteams": TeamworkAccountConfigurationSupportedClient_SkypeDefaultAndTeams,
		"skypeonly":            TeamworkAccountConfigurationSupportedClient_SkypeOnly,
		"teamsdefaultandskype": TeamworkAccountConfigurationSupportedClient_TeamsDefaultAndSkype,
		"teamsonly":            TeamworkAccountConfigurationSupportedClient_TeamsOnly,
		"unknown":              TeamworkAccountConfigurationSupportedClient_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkAccountConfigurationSupportedClient(input)
	return &out, nil
}
