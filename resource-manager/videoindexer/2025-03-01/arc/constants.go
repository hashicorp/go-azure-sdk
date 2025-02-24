package arc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionType string

const (
	PermissionTypeContributor PermissionType = "Contributor"
	PermissionTypeReader      PermissionType = "Reader"
)

func PossibleValuesForPermissionType() []string {
	return []string{
		string(PermissionTypeContributor),
		string(PermissionTypeReader),
	}
}

func (s *PermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionType(input string) (*PermissionType, error) {
	vals := map[string]PermissionType{
		"contributor": PermissionTypeContributor,
		"reader":      PermissionTypeReader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionType(input)
	return &out, nil
}

type Scope string

const (
	ScopeAccount Scope = "Account"
	ScopeVideo   Scope = "Video"
)

func PossibleValuesForScope() []string {
	return []string{
		string(ScopeAccount),
		string(ScopeVideo),
	}
}

func (s *Scope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScope(input string) (*Scope, error) {
	vals := map[string]Scope{
		"account": ScopeAccount,
		"video":   ScopeVideo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Scope(input)
	return &out, nil
}
