package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinitionPublishingState string

const (
	TeamsAppDefinitionPublishingState_Published TeamsAppDefinitionPublishingState = "published"
	TeamsAppDefinitionPublishingState_Rejected  TeamsAppDefinitionPublishingState = "rejected"
	TeamsAppDefinitionPublishingState_Submitted TeamsAppDefinitionPublishingState = "submitted"
)

func PossibleValuesForTeamsAppDefinitionPublishingState() []string {
	return []string{
		string(TeamsAppDefinitionPublishingState_Published),
		string(TeamsAppDefinitionPublishingState_Rejected),
		string(TeamsAppDefinitionPublishingState_Submitted),
	}
}

func (s *TeamsAppDefinitionPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDefinitionPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDefinitionPublishingState(input string) (*TeamsAppDefinitionPublishingState, error) {
	vals := map[string]TeamsAppDefinitionPublishingState{
		"published": TeamsAppDefinitionPublishingState_Published,
		"rejected":  TeamsAppDefinitionPublishingState_Rejected,
		"submitted": TeamsAppDefinitionPublishingState_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDefinitionPublishingState(input)
	return &out, nil
}
