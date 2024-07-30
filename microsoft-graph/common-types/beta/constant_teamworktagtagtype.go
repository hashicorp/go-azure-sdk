package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkTagTagType string

const (
	TeamworkTagTagType_Standard TeamworkTagTagType = "standard"
)

func PossibleValuesForTeamworkTagTagType() []string {
	return []string{
		string(TeamworkTagTagType_Standard),
	}
}

func (s *TeamworkTagTagType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkTagTagType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkTagTagType(input string) (*TeamworkTagTagType, error) {
	vals := map[string]TeamworkTagTagType{
		"standard": TeamworkTagTagType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkTagTagType(input)
	return &out, nil
}
