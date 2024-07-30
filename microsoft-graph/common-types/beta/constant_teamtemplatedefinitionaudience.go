package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateDefinitionAudience string

const (
	TeamTemplateDefinitionAudience_Organization TeamTemplateDefinitionAudience = "organization"
	TeamTemplateDefinitionAudience_Public       TeamTemplateDefinitionAudience = "public"
	TeamTemplateDefinitionAudience_User         TeamTemplateDefinitionAudience = "user"
)

func PossibleValuesForTeamTemplateDefinitionAudience() []string {
	return []string{
		string(TeamTemplateDefinitionAudience_Organization),
		string(TeamTemplateDefinitionAudience_Public),
		string(TeamTemplateDefinitionAudience_User),
	}
}

func (s *TeamTemplateDefinitionAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamTemplateDefinitionAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamTemplateDefinitionAudience(input string) (*TeamTemplateDefinitionAudience, error) {
	vals := map[string]TeamTemplateDefinitionAudience{
		"organization": TeamTemplateDefinitionAudience_Organization,
		"public":       TeamTemplateDefinitionAudience_Public,
		"user":         TeamTemplateDefinitionAudience_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamTemplateDefinitionAudience(input)
	return &out, nil
}
