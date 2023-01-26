package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmRollingUpgradeMonitoringPolicy struct {
	FailureAction             *ArmUpgradeFailureAction `json:"failureAction,omitempty"`
	HealthCheckRetryTimeout   *string                  `json:"healthCheckRetryTimeout,omitempty"`
	HealthCheckStableDuration *string                  `json:"healthCheckStableDuration,omitempty"`
	HealthCheckWaitDuration   *string                  `json:"healthCheckWaitDuration,omitempty"`
	UpgradeDomainTimeout      *string                  `json:"upgradeDomainTimeout,omitempty"`
	UpgradeTimeout            *string                  `json:"upgradeTimeout,omitempty"`
}
