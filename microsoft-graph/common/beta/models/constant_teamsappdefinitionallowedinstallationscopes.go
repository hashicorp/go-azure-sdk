package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinitionAllowedInstallationScopes string

const (
	TeamsAppDefinitionAllowedInstallationScopesgroupChat TeamsAppDefinitionAllowedInstallationScopes = "GroupChat"
	TeamsAppDefinitionAllowedInstallationScopespersonal  TeamsAppDefinitionAllowedInstallationScopes = "Personal"
	TeamsAppDefinitionAllowedInstallationScopesteam      TeamsAppDefinitionAllowedInstallationScopes = "Team"
)

func PossibleValuesForTeamsAppDefinitionAllowedInstallationScopes() []string {
	return []string{
		string(TeamsAppDefinitionAllowedInstallationScopesgroupChat),
		string(TeamsAppDefinitionAllowedInstallationScopespersonal),
		string(TeamsAppDefinitionAllowedInstallationScopesteam),
	}
}

func parseTeamsAppDefinitionAllowedInstallationScopes(input string) (*TeamsAppDefinitionAllowedInstallationScopes, error) {
	vals := map[string]TeamsAppDefinitionAllowedInstallationScopes{
		"groupchat": TeamsAppDefinitionAllowedInstallationScopesgroupChat,
		"personal":  TeamsAppDefinitionAllowedInstallationScopespersonal,
		"team":      TeamsAppDefinitionAllowedInstallationScopesteam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDefinitionAllowedInstallationScopes(input)
	return &out, nil
}
