package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LobbyBypassSettingsScope string

const (
	LobbyBypassSettingsScopeeveryone                    LobbyBypassSettingsScope = "Everyone"
	LobbyBypassSettingsScopeinvited                     LobbyBypassSettingsScope = "Invited"
	LobbyBypassSettingsScopeorganization                LobbyBypassSettingsScope = "Organization"
	LobbyBypassSettingsScopeorganizationAndFederated    LobbyBypassSettingsScope = "OrganizationAndFederated"
	LobbyBypassSettingsScopeorganizationExcludingGuests LobbyBypassSettingsScope = "OrganizationExcludingGuests"
	LobbyBypassSettingsScopeorganizer                   LobbyBypassSettingsScope = "Organizer"
)

func PossibleValuesForLobbyBypassSettingsScope() []string {
	return []string{
		string(LobbyBypassSettingsScopeeveryone),
		string(LobbyBypassSettingsScopeinvited),
		string(LobbyBypassSettingsScopeorganization),
		string(LobbyBypassSettingsScopeorganizationAndFederated),
		string(LobbyBypassSettingsScopeorganizationExcludingGuests),
		string(LobbyBypassSettingsScopeorganizer),
	}
}

func parseLobbyBypassSettingsScope(input string) (*LobbyBypassSettingsScope, error) {
	vals := map[string]LobbyBypassSettingsScope{
		"everyone":                    LobbyBypassSettingsScopeeveryone,
		"invited":                     LobbyBypassSettingsScopeinvited,
		"organization":                LobbyBypassSettingsScopeorganization,
		"organizationandfederated":    LobbyBypassSettingsScopeorganizationAndFederated,
		"organizationexcludingguests": LobbyBypassSettingsScopeorganizationExcludingGuests,
		"organizer":                   LobbyBypassSettingsScopeorganizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LobbyBypassSettingsScope(input)
	return &out, nil
}
