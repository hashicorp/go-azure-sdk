package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateDefinitionAudience string

const (
	TeamTemplateDefinitionAudienceorganization TeamTemplateDefinitionAudience = "Organization"
	TeamTemplateDefinitionAudiencepublic       TeamTemplateDefinitionAudience = "Public"
	TeamTemplateDefinitionAudienceuser         TeamTemplateDefinitionAudience = "User"
)

func PossibleValuesForTeamTemplateDefinitionAudience() []string {
	return []string{
		string(TeamTemplateDefinitionAudienceorganization),
		string(TeamTemplateDefinitionAudiencepublic),
		string(TeamTemplateDefinitionAudienceuser),
	}
}

func parseTeamTemplateDefinitionAudience(input string) (*TeamTemplateDefinitionAudience, error) {
	vals := map[string]TeamTemplateDefinitionAudience{
		"organization": TeamTemplateDefinitionAudienceorganization,
		"public":       TeamTemplateDefinitionAudiencepublic,
		"user":         TeamTemplateDefinitionAudienceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamTemplateDefinitionAudience(input)
	return &out, nil
}
