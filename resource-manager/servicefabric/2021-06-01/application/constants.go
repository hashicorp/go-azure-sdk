package application

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
