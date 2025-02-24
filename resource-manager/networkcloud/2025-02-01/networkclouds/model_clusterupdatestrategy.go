package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterUpdateStrategy struct {
	MaxUnavailable  *int64                    `json:"maxUnavailable,omitempty"`
	StrategyType    ClusterUpdateStrategyType `json:"strategyType"`
	ThresholdType   ValidationThresholdType   `json:"thresholdType"`
	ThresholdValue  int64                     `json:"thresholdValue"`
	WaitTimeMinutes *int64                    `json:"waitTimeMinutes,omitempty"`
}
