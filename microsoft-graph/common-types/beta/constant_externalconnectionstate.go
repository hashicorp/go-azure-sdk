package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectionState string

const (
	ExternalConnectionState_Draft         ExternalConnectionState = "draft"
	ExternalConnectionState_LimitExceeded ExternalConnectionState = "limitExceeded"
	ExternalConnectionState_Obsolete      ExternalConnectionState = "obsolete"
	ExternalConnectionState_Ready         ExternalConnectionState = "ready"
)

func PossibleValuesForExternalConnectionState() []string {
	return []string{
		string(ExternalConnectionState_Draft),
		string(ExternalConnectionState_LimitExceeded),
		string(ExternalConnectionState_Obsolete),
		string(ExternalConnectionState_Ready),
	}
}

func (s *ExternalConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectionState(input string) (*ExternalConnectionState, error) {
	vals := map[string]ExternalConnectionState{
		"draft":         ExternalConnectionState_Draft,
		"limitexceeded": ExternalConnectionState_LimitExceeded,
		"obsolete":      ExternalConnectionState_Obsolete,
		"ready":         ExternalConnectionState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectionState(input)
	return &out, nil
}
