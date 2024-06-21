package adaptivenetworkhardenings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveNetworkHardeningProperties struct {
	EffectiveNetworkSecurityGroups *[]EffectiveNetworkSecurityGroups `json:"effectiveNetworkSecurityGroups,omitempty"`
	Rules                          *[]Rule                           `json:"rules,omitempty"`
	RulesCalculationTime           *string                           `json:"rulesCalculationTime,omitempty"`
}
