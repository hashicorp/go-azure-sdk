package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExcludeTargetTargetType string

const (
	ExcludeTargetTargetType_Group ExcludeTargetTargetType = "group"
	ExcludeTargetTargetType_User  ExcludeTargetTargetType = "user"
)

func PossibleValuesForExcludeTargetTargetType() []string {
	return []string{
		string(ExcludeTargetTargetType_Group),
		string(ExcludeTargetTargetType_User),
	}
}

func (s *ExcludeTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExcludeTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExcludeTargetTargetType(input string) (*ExcludeTargetTargetType, error) {
	vals := map[string]ExcludeTargetTargetType{
		"group": ExcludeTargetTargetType_Group,
		"user":  ExcludeTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExcludeTargetTargetType(input)
	return &out, nil
}
