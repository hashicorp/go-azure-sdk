package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppDependencyTargetType string

const (
	MobileAppDependencyTargetType_Child  MobileAppDependencyTargetType = "child"
	MobileAppDependencyTargetType_Parent MobileAppDependencyTargetType = "parent"
)

func PossibleValuesForMobileAppDependencyTargetType() []string {
	return []string{
		string(MobileAppDependencyTargetType_Child),
		string(MobileAppDependencyTargetType_Parent),
	}
}

func (s *MobileAppDependencyTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppDependencyTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppDependencyTargetType(input string) (*MobileAppDependencyTargetType, error) {
	vals := map[string]MobileAppDependencyTargetType{
		"child":  MobileAppDependencyTargetType_Child,
		"parent": MobileAppDependencyTargetType_Parent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppDependencyTargetType(input)
	return &out, nil
}
