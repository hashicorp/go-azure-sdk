package vaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpgradeDetails struct {
	EndTimeUtc         *string            `json:"endTimeUtc,omitempty"`
	LastUpdatedTimeUtc *string            `json:"lastUpdatedTimeUtc,omitempty"`
	Message            *string            `json:"message,omitempty"`
	OperationId        *string            `json:"operationId,omitempty"`
	PreviousResourceId *string            `json:"previousResourceId,omitempty"`
	StartTimeUtc       *string            `json:"startTimeUtc,omitempty"`
	Status             *VaultUpgradeState `json:"status,omitempty"`
	TriggerType        *TriggerType       `json:"triggerType,omitempty"`
	UpgradedResourceId *string            `json:"upgradedResourceId,omitempty"`
}
