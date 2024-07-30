package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppRelationshipTargetType string

const (
	MobileAppRelationshipTargetType_Child  MobileAppRelationshipTargetType = "child"
	MobileAppRelationshipTargetType_Parent MobileAppRelationshipTargetType = "parent"
)

func PossibleValuesForMobileAppRelationshipTargetType() []string {
	return []string{
		string(MobileAppRelationshipTargetType_Child),
		string(MobileAppRelationshipTargetType_Parent),
	}
}

func (s *MobileAppRelationshipTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppRelationshipTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppRelationshipTargetType(input string) (*MobileAppRelationshipTargetType, error) {
	vals := map[string]MobileAppRelationshipTargetType{
		"child":  MobileAppRelationshipTargetType_Child,
		"parent": MobileAppRelationshipTargetType_Parent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppRelationshipTargetType(input)
	return &out, nil
}
