package application

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmUpgradeFailureAction string

const (
	ArmUpgradeFailureActionManual   ArmUpgradeFailureAction = "Manual"
	ArmUpgradeFailureActionRollback ArmUpgradeFailureAction = "Rollback"
)

func PossibleValuesForArmUpgradeFailureAction() []string {
	return []string{
		string(ArmUpgradeFailureActionManual),
		string(ArmUpgradeFailureActionRollback),
	}
}

func parseArmUpgradeFailureAction(input string) (*ArmUpgradeFailureAction, error) {
	vals := map[string]ArmUpgradeFailureAction{
		"manual":   ArmUpgradeFailureActionManual,
		"rollback": ArmUpgradeFailureActionRollback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArmUpgradeFailureAction(input)
	return &out, nil
}

type RollingUpgradeMode string

const (
	RollingUpgradeModeInvalid           RollingUpgradeMode = "Invalid"
	RollingUpgradeModeMonitored         RollingUpgradeMode = "Monitored"
	RollingUpgradeModeUnmonitoredAuto   RollingUpgradeMode = "UnmonitoredAuto"
	RollingUpgradeModeUnmonitoredManual RollingUpgradeMode = "UnmonitoredManual"
)

func PossibleValuesForRollingUpgradeMode() []string {
	return []string{
		string(RollingUpgradeModeInvalid),
		string(RollingUpgradeModeMonitored),
		string(RollingUpgradeModeUnmonitoredAuto),
		string(RollingUpgradeModeUnmonitoredManual),
	}
}

func parseRollingUpgradeMode(input string) (*RollingUpgradeMode, error) {
	vals := map[string]RollingUpgradeMode{
		"invalid":           RollingUpgradeModeInvalid,
		"monitored":         RollingUpgradeModeMonitored,
		"unmonitoredauto":   RollingUpgradeModeUnmonitoredAuto,
		"unmonitoredmanual": RollingUpgradeModeUnmonitoredManual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RollingUpgradeMode(input)
	return &out, nil
}
