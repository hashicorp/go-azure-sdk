package agentapplication

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrafficRoutingRule struct {
	DeploymentId      *string `json:"deploymentId,omitempty"`
	Description       *string `json:"description,omitempty"`
	RuleId            *string `json:"ruleId,omitempty"`
	TrafficPercentage *int64  `json:"trafficPercentage,omitempty"`
}
