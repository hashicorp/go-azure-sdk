package postrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleCounter struct {
	AppSeen              *AppSeenData `json:"appSeen,omitempty"`
	FirewallName         *string      `json:"firewallName,omitempty"`
	HitCount             *int64       `json:"hitCount,omitempty"`
	LastUpdatedTimestamp *string      `json:"lastUpdatedTimestamp,omitempty"`
	Priority             string       `json:"priority"`
	RequestTimestamp     *string      `json:"requestTimestamp,omitempty"`
	RuleListName         *string      `json:"ruleListName,omitempty"`
	RuleName             string       `json:"ruleName"`
	RuleStackName        *string      `json:"ruleStackName,omitempty"`
	Timestamp            *string      `json:"timestamp,omitempty"`
}
