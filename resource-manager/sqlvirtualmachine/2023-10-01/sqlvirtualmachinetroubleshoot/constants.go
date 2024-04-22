package sqlvirtualmachinetroubleshoot

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TroubleshootingScenario string

const (
	TroubleshootingScenarioUnhealthyReplica TroubleshootingScenario = "UnhealthyReplica"
)

func PossibleValuesForTroubleshootingScenario() []string {
	return []string{
		string(TroubleshootingScenarioUnhealthyReplica),
	}
}

func (s *TroubleshootingScenario) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTroubleshootingScenario(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTroubleshootingScenario(input string) (*TroubleshootingScenario, error) {
	vals := map[string]TroubleshootingScenario{
		"unhealthyreplica": TroubleshootingScenarioUnhealthyReplica,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TroubleshootingScenario(input)
	return &out, nil
}
