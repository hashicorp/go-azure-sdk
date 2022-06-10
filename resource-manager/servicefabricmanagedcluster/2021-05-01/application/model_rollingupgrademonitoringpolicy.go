package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RollingUpgradeMonitoringPolicy struct {
	FailureAction             FailureAction `json:"failureAction"`
	HealthCheckRetryTimeout   string        `json:"healthCheckRetryTimeout"`
	HealthCheckStableDuration string        `json:"healthCheckStableDuration"`
	HealthCheckWaitDuration   string        `json:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      string        `json:"upgradeDomainTimeout"`
	UpgradeTimeout            string        `json:"upgradeTimeout"`
}
