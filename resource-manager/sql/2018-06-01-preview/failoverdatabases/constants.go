package failoverdatabases

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicaType string

const (
	ReplicaTypePrimary           ReplicaType = "Primary"
	ReplicaTypeReadableSecondary ReplicaType = "ReadableSecondary"
)

func PossibleValuesForReplicaType() []string {
	return []string{
		string(ReplicaTypePrimary),
		string(ReplicaTypeReadableSecondary),
	}
}

func (s *ReplicaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicaType(input string) (*ReplicaType, error) {
	vals := map[string]ReplicaType{
		"primary":           ReplicaTypePrimary,
		"readablesecondary": ReplicaTypeReadableSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicaType(input)
	return &out, nil
}
