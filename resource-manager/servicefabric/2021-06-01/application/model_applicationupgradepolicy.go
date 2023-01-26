package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationUpgradePolicy struct {
	ApplicationHealthPolicy        *ArmApplicationHealthPolicy        `json:"applicationHealthPolicy,omitempty"`
	ForceRestart                   *bool                              `json:"forceRestart,omitempty"`
	RecreateApplication            *bool                              `json:"recreateApplication,omitempty"`
	RollingUpgradeMonitoringPolicy *ArmRollingUpgradeMonitoringPolicy `json:"rollingUpgradeMonitoringPolicy,omitempty"`
	UpgradeMode                    *RollingUpgradeMode                `json:"upgradeMode,omitempty"`
	UpgradeReplicaSetCheckTimeout  *string                            `json:"upgradeReplicaSetCheckTimeout,omitempty"`
}
