package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersistentBrowserSessionControlMode string

const (
	PersistentBrowserSessionControlMode_Always PersistentBrowserSessionControlMode = "always"
	PersistentBrowserSessionControlMode_Never  PersistentBrowserSessionControlMode = "never"
)

func PossibleValuesForPersistentBrowserSessionControlMode() []string {
	return []string{
		string(PersistentBrowserSessionControlMode_Always),
		string(PersistentBrowserSessionControlMode_Never),
	}
}

func (s *PersistentBrowserSessionControlMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersistentBrowserSessionControlMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersistentBrowserSessionControlMode(input string) (*PersistentBrowserSessionControlMode, error) {
	vals := map[string]PersistentBrowserSessionControlMode{
		"always": PersistentBrowserSessionControlMode_Always,
		"never":  PersistentBrowserSessionControlMode_Never,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersistentBrowserSessionControlMode(input)
	return &out, nil
}
