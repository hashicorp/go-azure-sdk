package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnectionState string

const (
	ExternalConnectorsExternalConnectionState_Draft         ExternalConnectorsExternalConnectionState = "draft"
	ExternalConnectorsExternalConnectionState_LimitExceeded ExternalConnectorsExternalConnectionState = "limitExceeded"
	ExternalConnectorsExternalConnectionState_Obsolete      ExternalConnectorsExternalConnectionState = "obsolete"
	ExternalConnectorsExternalConnectionState_Ready         ExternalConnectorsExternalConnectionState = "ready"
)

func PossibleValuesForExternalConnectorsExternalConnectionState() []string {
	return []string{
		string(ExternalConnectorsExternalConnectionState_Draft),
		string(ExternalConnectorsExternalConnectionState_LimitExceeded),
		string(ExternalConnectorsExternalConnectionState_Obsolete),
		string(ExternalConnectorsExternalConnectionState_Ready),
	}
}

func (s *ExternalConnectorsExternalConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsExternalConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsExternalConnectionState(input string) (*ExternalConnectorsExternalConnectionState, error) {
	vals := map[string]ExternalConnectorsExternalConnectionState{
		"draft":         ExternalConnectorsExternalConnectionState_Draft,
		"limitexceeded": ExternalConnectorsExternalConnectionState_LimitExceeded,
		"obsolete":      ExternalConnectorsExternalConnectionState_Obsolete,
		"ready":         ExternalConnectorsExternalConnectionState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalConnectionState(input)
	return &out, nil
}
