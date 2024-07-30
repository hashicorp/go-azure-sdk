package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LobbyBypassSettingsScope string

const (
	LobbyBypassSettingsScope_Everyone                    LobbyBypassSettingsScope = "everyone"
	LobbyBypassSettingsScope_Invited                     LobbyBypassSettingsScope = "invited"
	LobbyBypassSettingsScope_Organization                LobbyBypassSettingsScope = "organization"
	LobbyBypassSettingsScope_OrganizationAndFederated    LobbyBypassSettingsScope = "organizationAndFederated"
	LobbyBypassSettingsScope_OrganizationExcludingGuests LobbyBypassSettingsScope = "organizationExcludingGuests"
	LobbyBypassSettingsScope_Organizer                   LobbyBypassSettingsScope = "organizer"
)

func PossibleValuesForLobbyBypassSettingsScope() []string {
	return []string{
		string(LobbyBypassSettingsScope_Everyone),
		string(LobbyBypassSettingsScope_Invited),
		string(LobbyBypassSettingsScope_Organization),
		string(LobbyBypassSettingsScope_OrganizationAndFederated),
		string(LobbyBypassSettingsScope_OrganizationExcludingGuests),
		string(LobbyBypassSettingsScope_Organizer),
	}
}

func (s *LobbyBypassSettingsScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLobbyBypassSettingsScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLobbyBypassSettingsScope(input string) (*LobbyBypassSettingsScope, error) {
	vals := map[string]LobbyBypassSettingsScope{
		"everyone":                    LobbyBypassSettingsScope_Everyone,
		"invited":                     LobbyBypassSettingsScope_Invited,
		"organization":                LobbyBypassSettingsScope_Organization,
		"organizationandfederated":    LobbyBypassSettingsScope_OrganizationAndFederated,
		"organizationexcludingguests": LobbyBypassSettingsScope_OrganizationExcludingGuests,
		"organizer":                   LobbyBypassSettingsScope_Organizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LobbyBypassSettingsScope(input)
	return &out, nil
}
