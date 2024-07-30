package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeTagTargetType string

const (
	ScopeTagTargetType_Device ScopeTagTargetType = "device"
	ScopeTagTargetType_None   ScopeTagTargetType = "none"
	ScopeTagTargetType_User   ScopeTagTargetType = "user"
)

func PossibleValuesForScopeTagTargetType() []string {
	return []string{
		string(ScopeTagTargetType_Device),
		string(ScopeTagTargetType_None),
		string(ScopeTagTargetType_User),
	}
}

func (s *ScopeTagTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScopeTagTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScopeTagTargetType(input string) (*ScopeTagTargetType, error) {
	vals := map[string]ScopeTagTargetType{
		"device": ScopeTagTargetType_Device,
		"none":   ScopeTagTargetType_None,
		"user":   ScopeTagTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScopeTagTargetType(input)
	return &out, nil
}
