package cdnwebapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomRule struct {
	Action          ActionType              `json:"action"`
	EnabledState    *CustomRuleEnabledState `json:"enabledState,omitempty"`
	MatchConditions []MatchCondition        `json:"matchConditions"`
	Name            string                  `json:"name"`
	Priority        int64                   `json:"priority"`
}
