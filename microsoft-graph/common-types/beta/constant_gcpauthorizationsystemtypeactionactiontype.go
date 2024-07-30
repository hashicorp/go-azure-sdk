package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpAuthorizationSystemTypeActionActionType string

const (
	GcpAuthorizationSystemTypeActionActionType_Delete GcpAuthorizationSystemTypeActionActionType = "delete"
	GcpAuthorizationSystemTypeActionActionType_Read   GcpAuthorizationSystemTypeActionActionType = "read"
)

func PossibleValuesForGcpAuthorizationSystemTypeActionActionType() []string {
	return []string{
		string(GcpAuthorizationSystemTypeActionActionType_Delete),
		string(GcpAuthorizationSystemTypeActionActionType_Read),
	}
}

func (s *GcpAuthorizationSystemTypeActionActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGcpAuthorizationSystemTypeActionActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGcpAuthorizationSystemTypeActionActionType(input string) (*GcpAuthorizationSystemTypeActionActionType, error) {
	vals := map[string]GcpAuthorizationSystemTypeActionActionType{
		"delete": GcpAuthorizationSystemTypeActionActionType_Delete,
		"read":   GcpAuthorizationSystemTypeActionActionType_Read,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GcpAuthorizationSystemTypeActionActionType(input)
	return &out, nil
}
