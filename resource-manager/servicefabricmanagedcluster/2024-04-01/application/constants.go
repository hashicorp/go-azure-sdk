package application

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailureAction string

const (
	FailureActionManual   FailureAction = "Manual"
	FailureActionRollback FailureAction = "Rollback"
)

func PossibleValuesForFailureAction() []string {
	return []string{
		string(FailureActionManual),
		string(FailureActionRollback),
	}
}

func (s *FailureAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFailureAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFailureAction(input string) (*FailureAction, error) {
	vals := map[string]FailureAction{
		"manual":   FailureActionManual,
		"rollback": FailureActionRollback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FailureAction(input)
	return &out, nil
}

type RollingUpgradeMode string

const (
	RollingUpgradeModeMonitored       RollingUpgradeMode = "Monitored"
	RollingUpgradeModeUnmonitoredAuto RollingUpgradeMode = "UnmonitoredAuto"
)

func PossibleValuesForRollingUpgradeMode() []string {
	return []string{
		string(RollingUpgradeModeMonitored),
		string(RollingUpgradeModeUnmonitoredAuto),
	}
}

func (s *RollingUpgradeMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRollingUpgradeMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRollingUpgradeMode(input string) (*RollingUpgradeMode, error) {
	vals := map[string]RollingUpgradeMode{
		"monitored":       RollingUpgradeModeMonitored,
		"unmonitoredauto": RollingUpgradeModeUnmonitoredAuto,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RollingUpgradeMode(input)
	return &out, nil
}
