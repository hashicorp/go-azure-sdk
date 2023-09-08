package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleAction string

const (
	NetworkaccessM365ForwardingRuleActionbypass  NetworkaccessM365ForwardingRuleAction = "Bypass"
	NetworkaccessM365ForwardingRuleActionforward NetworkaccessM365ForwardingRuleAction = "Forward"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleAction() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleActionbypass),
		string(NetworkaccessM365ForwardingRuleActionforward),
	}
}

func parseNetworkaccessM365ForwardingRuleAction(input string) (*NetworkaccessM365ForwardingRuleAction, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleAction{
		"bypass":  NetworkaccessM365ForwardingRuleActionbypass,
		"forward": NetworkaccessM365ForwardingRuleActionforward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleAction(input)
	return &out, nil
}
