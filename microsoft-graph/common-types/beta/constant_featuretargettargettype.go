package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureTargetTargetType string

const (
	FeatureTargetTargetType_AdministrativeUnit FeatureTargetTargetType = "administrativeUnit"
	FeatureTargetTargetType_Group              FeatureTargetTargetType = "group"
	FeatureTargetTargetType_Role               FeatureTargetTargetType = "role"
)

func PossibleValuesForFeatureTargetTargetType() []string {
	return []string{
		string(FeatureTargetTargetType_AdministrativeUnit),
		string(FeatureTargetTargetType_Group),
		string(FeatureTargetTargetType_Role),
	}
}

func (s *FeatureTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureTargetTargetType(input string) (*FeatureTargetTargetType, error) {
	vals := map[string]FeatureTargetTargetType{
		"administrativeunit": FeatureTargetTargetType_AdministrativeUnit,
		"group":              FeatureTargetTargetType_Group,
		"role":               FeatureTargetTargetType_Role,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureTargetTargetType(input)
	return &out, nil
}
