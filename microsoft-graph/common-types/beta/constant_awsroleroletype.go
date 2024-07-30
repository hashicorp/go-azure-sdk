package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsRoleRoleType string

const (
	AwsRoleRoleType_Custom AwsRoleRoleType = "custom"
	AwsRoleRoleType_System AwsRoleRoleType = "system"
)

func PossibleValuesForAwsRoleRoleType() []string {
	return []string{
		string(AwsRoleRoleType_Custom),
		string(AwsRoleRoleType_System),
	}
}

func (s *AwsRoleRoleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsRoleRoleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsRoleRoleType(input string) (*AwsRoleRoleType, error) {
	vals := map[string]AwsRoleRoleType{
		"custom": AwsRoleRoleType_Custom,
		"system": AwsRoleRoleType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsRoleRoleType(input)
	return &out, nil
}
