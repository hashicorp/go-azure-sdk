package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentStatus string

const (
	OnPremisesAgentStatus_Active   OnPremisesAgentStatus = "active"
	OnPremisesAgentStatus_Inactive OnPremisesAgentStatus = "inactive"
)

func PossibleValuesForOnPremisesAgentStatus() []string {
	return []string{
		string(OnPremisesAgentStatus_Active),
		string(OnPremisesAgentStatus_Inactive),
	}
}

func (s *OnPremisesAgentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesAgentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesAgentStatus(input string) (*OnPremisesAgentStatus, error) {
	vals := map[string]OnPremisesAgentStatus{
		"active":   OnPremisesAgentStatus_Active,
		"inactive": OnPremisesAgentStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAgentStatus(input)
	return &out, nil
}
