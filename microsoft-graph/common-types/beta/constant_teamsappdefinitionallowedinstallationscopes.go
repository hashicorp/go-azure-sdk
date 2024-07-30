package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinitionAllowedInstallationScopes string

const (
	TeamsAppDefinitionAllowedInstallationScopes_GroupChat TeamsAppDefinitionAllowedInstallationScopes = "groupChat"
	TeamsAppDefinitionAllowedInstallationScopes_Personal  TeamsAppDefinitionAllowedInstallationScopes = "personal"
	TeamsAppDefinitionAllowedInstallationScopes_Team      TeamsAppDefinitionAllowedInstallationScopes = "team"
)

func PossibleValuesForTeamsAppDefinitionAllowedInstallationScopes() []string {
	return []string{
		string(TeamsAppDefinitionAllowedInstallationScopes_GroupChat),
		string(TeamsAppDefinitionAllowedInstallationScopes_Personal),
		string(TeamsAppDefinitionAllowedInstallationScopes_Team),
	}
}

func (s *TeamsAppDefinitionAllowedInstallationScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDefinitionAllowedInstallationScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDefinitionAllowedInstallationScopes(input string) (*TeamsAppDefinitionAllowedInstallationScopes, error) {
	vals := map[string]TeamsAppDefinitionAllowedInstallationScopes{
		"groupchat": TeamsAppDefinitionAllowedInstallationScopes_GroupChat,
		"personal":  TeamsAppDefinitionAllowedInstallationScopes_Personal,
		"team":      TeamsAppDefinitionAllowedInstallationScopes_Team,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDefinitionAllowedInstallationScopes(input)
	return &out, nil
}
