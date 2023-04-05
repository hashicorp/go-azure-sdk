package application

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
