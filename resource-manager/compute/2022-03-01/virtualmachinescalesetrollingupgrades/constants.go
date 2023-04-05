package virtualmachinescalesetrollingupgrades

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RollingUpgradeActionType string

const (
	RollingUpgradeActionTypeCancel RollingUpgradeActionType = "Cancel"
	RollingUpgradeActionTypeStart  RollingUpgradeActionType = "Start"
)

func PossibleValuesForRollingUpgradeActionType() []string {
	return []string{
		string(RollingUpgradeActionTypeCancel),
		string(RollingUpgradeActionTypeStart),
	}
}

type RollingUpgradeStatusCode string

const (
	RollingUpgradeStatusCodeCancelled      RollingUpgradeStatusCode = "Cancelled"
	RollingUpgradeStatusCodeCompleted      RollingUpgradeStatusCode = "Completed"
	RollingUpgradeStatusCodeFaulted        RollingUpgradeStatusCode = "Faulted"
	RollingUpgradeStatusCodeRollingForward RollingUpgradeStatusCode = "RollingForward"
)

func PossibleValuesForRollingUpgradeStatusCode() []string {
	return []string{
		string(RollingUpgradeStatusCodeCancelled),
		string(RollingUpgradeStatusCodeCompleted),
		string(RollingUpgradeStatusCodeFaulted),
		string(RollingUpgradeStatusCodeRollingForward),
	}
}
