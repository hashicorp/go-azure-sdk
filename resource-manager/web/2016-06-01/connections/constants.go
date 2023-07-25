package connections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkState string

const (
	LinkStateAuthenticated   LinkState = "Authenticated"
	LinkStateError           LinkState = "Error"
	LinkStateUnauthenticated LinkState = "Unauthenticated"
)

func PossibleValuesForLinkState() []string {
	return []string{
		string(LinkStateAuthenticated),
		string(LinkStateError),
		string(LinkStateUnauthenticated),
	}
}

func (s *LinkState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLinkState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLinkState(input string) (*LinkState, error) {
	vals := map[string]LinkState{
		"authenticated":   LinkStateAuthenticated,
		"error":           LinkStateError,
		"unauthenticated": LinkStateUnauthenticated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LinkState(input)
	return &out, nil
}
