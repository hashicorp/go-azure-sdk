package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppResourceSpecificPermissionPermissionType string

const (
	TeamsAppResourceSpecificPermissionPermissionType_Application TeamsAppResourceSpecificPermissionPermissionType = "application"
	TeamsAppResourceSpecificPermissionPermissionType_Delegated   TeamsAppResourceSpecificPermissionPermissionType = "delegated"
)

func PossibleValuesForTeamsAppResourceSpecificPermissionPermissionType() []string {
	return []string{
		string(TeamsAppResourceSpecificPermissionPermissionType_Application),
		string(TeamsAppResourceSpecificPermissionPermissionType_Delegated),
	}
}

func (s *TeamsAppResourceSpecificPermissionPermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppResourceSpecificPermissionPermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppResourceSpecificPermissionPermissionType(input string) (*TeamsAppResourceSpecificPermissionPermissionType, error) {
	vals := map[string]TeamsAppResourceSpecificPermissionPermissionType{
		"application": TeamsAppResourceSpecificPermissionPermissionType_Application,
		"delegated":   TeamsAppResourceSpecificPermissionPermissionType_Delegated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppResourceSpecificPermissionPermissionType(input)
	return &out, nil
}
