package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleCategory string

const (
	NetworkaccessM365ForwardingRuleCategoryallow     NetworkaccessM365ForwardingRuleCategory = "Allow"
	NetworkaccessM365ForwardingRuleCategorydefault   NetworkaccessM365ForwardingRuleCategory = "Default"
	NetworkaccessM365ForwardingRuleCategoryoptimized NetworkaccessM365ForwardingRuleCategory = "Optimized"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleCategory() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleCategoryallow),
		string(NetworkaccessM365ForwardingRuleCategorydefault),
		string(NetworkaccessM365ForwardingRuleCategoryoptimized),
	}
}

func parseNetworkaccessM365ForwardingRuleCategory(input string) (*NetworkaccessM365ForwardingRuleCategory, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleCategory{
		"allow":     NetworkaccessM365ForwardingRuleCategoryallow,
		"default":   NetworkaccessM365ForwardingRuleCategorydefault,
		"optimized": NetworkaccessM365ForwardingRuleCategoryoptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleCategory(input)
	return &out, nil
}
