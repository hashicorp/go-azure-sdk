package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupersedenceTargetType string

const (
	MobileAppSupersedenceTargetType_Child  MobileAppSupersedenceTargetType = "child"
	MobileAppSupersedenceTargetType_Parent MobileAppSupersedenceTargetType = "parent"
)

func PossibleValuesForMobileAppSupersedenceTargetType() []string {
	return []string{
		string(MobileAppSupersedenceTargetType_Child),
		string(MobileAppSupersedenceTargetType_Parent),
	}
}

func (s *MobileAppSupersedenceTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppSupersedenceTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppSupersedenceTargetType(input string) (*MobileAppSupersedenceTargetType, error) {
	vals := map[string]MobileAppSupersedenceTargetType{
		"child":  MobileAppSupersedenceTargetType_Child,
		"parent": MobileAppSupersedenceTargetType_Parent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupersedenceTargetType(input)
	return &out, nil
}
