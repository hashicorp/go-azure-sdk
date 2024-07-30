package joinedteaminstalledapp

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppResourceSpecificPermissionPermissionType string

const (
	TeamsAppResourceSpecificPermissionPermissionTypeApplication TeamsAppResourceSpecificPermissionPermissionType = "application"
	TeamsAppResourceSpecificPermissionPermissionTypeDelegated   TeamsAppResourceSpecificPermissionPermissionType = "delegated"
)

func PossibleValuesForTeamsAppResourceSpecificPermissionPermissionType() []string {
	return []string{
		string(TeamsAppResourceSpecificPermissionPermissionTypeApplication),
		string(TeamsAppResourceSpecificPermissionPermissionTypeDelegated),
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
		"application": TeamsAppResourceSpecificPermissionPermissionTypeApplication,
		"delegated":   TeamsAppResourceSpecificPermissionPermissionTypeDelegated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppResourceSpecificPermissionPermissionType(input)
	return &out, nil
}
