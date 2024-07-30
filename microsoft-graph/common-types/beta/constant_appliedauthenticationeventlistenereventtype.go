package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedAuthenticationEventListenerEventType string

const (
	AppliedAuthenticationEventListenerEventType_PageRenderStart    AppliedAuthenticationEventListenerEventType = "pageRenderStart"
	AppliedAuthenticationEventListenerEventType_TokenIssuanceStart AppliedAuthenticationEventListenerEventType = "tokenIssuanceStart"
)

func PossibleValuesForAppliedAuthenticationEventListenerEventType() []string {
	return []string{
		string(AppliedAuthenticationEventListenerEventType_PageRenderStart),
		string(AppliedAuthenticationEventListenerEventType_TokenIssuanceStart),
	}
}

func (s *AppliedAuthenticationEventListenerEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppliedAuthenticationEventListenerEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppliedAuthenticationEventListenerEventType(input string) (*AppliedAuthenticationEventListenerEventType, error) {
	vals := map[string]AppliedAuthenticationEventListenerEventType{
		"pagerenderstart":    AppliedAuthenticationEventListenerEventType_PageRenderStart,
		"tokenissuancestart": AppliedAuthenticationEventListenerEventType_TokenIssuanceStart,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedAuthenticationEventListenerEventType(input)
	return &out, nil
}
