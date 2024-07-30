package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationOperationStatus string

const (
	AttackSimulationOperationStatus_Failed     AttackSimulationOperationStatus = "failed"
	AttackSimulationOperationStatus_NotStarted AttackSimulationOperationStatus = "notStarted"
	AttackSimulationOperationStatus_Running    AttackSimulationOperationStatus = "running"
	AttackSimulationOperationStatus_Succeeded  AttackSimulationOperationStatus = "succeeded"
)

func PossibleValuesForAttackSimulationOperationStatus() []string {
	return []string{
		string(AttackSimulationOperationStatus_Failed),
		string(AttackSimulationOperationStatus_NotStarted),
		string(AttackSimulationOperationStatus_Running),
		string(AttackSimulationOperationStatus_Succeeded),
	}
}

func (s *AttackSimulationOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttackSimulationOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttackSimulationOperationStatus(input string) (*AttackSimulationOperationStatus, error) {
	vals := map[string]AttackSimulationOperationStatus{
		"failed":     AttackSimulationOperationStatus_Failed,
		"notstarted": AttackSimulationOperationStatus_NotStarted,
		"running":    AttackSimulationOperationStatus_Running,
		"succeeded":  AttackSimulationOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackSimulationOperationStatus(input)
	return &out, nil
}
