package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamVisibility string

const (
	TeamVisibility_HiddenMembership TeamVisibility = "hiddenMembership"
	TeamVisibility_Private          TeamVisibility = "private"
	TeamVisibility_Public           TeamVisibility = "public"
)

func PossibleValuesForTeamVisibility() []string {
	return []string{
		string(TeamVisibility_HiddenMembership),
		string(TeamVisibility_Private),
		string(TeamVisibility_Public),
	}
}

func (s *TeamVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamVisibility(input string) (*TeamVisibility, error) {
	vals := map[string]TeamVisibility{
		"hiddenmembership": TeamVisibility_HiddenMembership,
		"private":          TeamVisibility_Private,
		"public":           TeamVisibility_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamVisibility(input)
	return &out, nil
}
