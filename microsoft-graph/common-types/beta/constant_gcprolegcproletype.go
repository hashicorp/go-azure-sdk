package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpRoleGcpRoleType string

const (
	GcpRoleGcpRoleType_Custom GcpRoleGcpRoleType = "custom"
	GcpRoleGcpRoleType_System GcpRoleGcpRoleType = "system"
)

func PossibleValuesForGcpRoleGcpRoleType() []string {
	return []string{
		string(GcpRoleGcpRoleType_Custom),
		string(GcpRoleGcpRoleType_System),
	}
}

func (s *GcpRoleGcpRoleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGcpRoleGcpRoleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGcpRoleGcpRoleType(input string) (*GcpRoleGcpRoleType, error) {
	vals := map[string]GcpRoleGcpRoleType{
		"custom": GcpRoleGcpRoleType_Custom,
		"system": GcpRoleGcpRoleType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GcpRoleGcpRoleType(input)
	return &out, nil
}
