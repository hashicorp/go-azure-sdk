package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludeTargetTargetType string

const (
	IncludeTargetTargetType_Group IncludeTargetTargetType = "group"
	IncludeTargetTargetType_User  IncludeTargetTargetType = "user"
)

func PossibleValuesForIncludeTargetTargetType() []string {
	return []string{
		string(IncludeTargetTargetType_Group),
		string(IncludeTargetTargetType_User),
	}
}

func (s *IncludeTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncludeTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncludeTargetTargetType(input string) (*IncludeTargetTargetType, error) {
	vals := map[string]IncludeTargetTargetType{
		"group": IncludeTargetTargetType_Group,
		"user":  IncludeTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncludeTargetTargetType(input)
	return &out, nil
}
