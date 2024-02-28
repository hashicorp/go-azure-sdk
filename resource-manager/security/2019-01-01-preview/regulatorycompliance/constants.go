package regulatorycompliance

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type State string

const (
	StateFailed      State = "Failed"
	StatePassed      State = "Passed"
	StateSkipped     State = "Skipped"
	StateUnsupported State = "Unsupported"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateFailed),
		string(StatePassed),
		string(StateSkipped),
		string(StateUnsupported),
	}
}

func (s *State) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseState(input string) (*State, error) {
	vals := map[string]State{
		"failed":      StateFailed,
		"passed":      StatePassed,
		"skipped":     StateSkipped,
		"unsupported": StateUnsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}
