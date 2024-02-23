package repositories

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RepositoryAccessKind string

const (
	RepositoryAccessKindApp   RepositoryAccessKind = "App"
	RepositoryAccessKindOAuth RepositoryAccessKind = "OAuth"
	RepositoryAccessKindPAT   RepositoryAccessKind = "PAT"
)

func PossibleValuesForRepositoryAccessKind() []string {
	return []string{
		string(RepositoryAccessKindApp),
		string(RepositoryAccessKindOAuth),
		string(RepositoryAccessKindPAT),
	}
}

func (s *RepositoryAccessKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRepositoryAccessKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRepositoryAccessKind(input string) (*RepositoryAccessKind, error) {
	vals := map[string]RepositoryAccessKind{
		"app":   RepositoryAccessKindApp,
		"oauth": RepositoryAccessKindOAuth,
		"pat":   RepositoryAccessKindPAT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RepositoryAccessKind(input)
	return &out, nil
}
