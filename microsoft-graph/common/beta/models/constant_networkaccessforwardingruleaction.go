package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingRuleAction string

const (
	NetworkaccessForwardingRuleActionbypass  NetworkaccessForwardingRuleAction = "Bypass"
	NetworkaccessForwardingRuleActionforward NetworkaccessForwardingRuleAction = "Forward"
)

func PossibleValuesForNetworkaccessForwardingRuleAction() []string {
	return []string{
		string(NetworkaccessForwardingRuleActionbypass),
		string(NetworkaccessForwardingRuleActionforward),
	}
}

func parseNetworkaccessForwardingRuleAction(input string) (*NetworkaccessForwardingRuleAction, error) {
	vals := map[string]NetworkaccessForwardingRuleAction{
		"bypass":  NetworkaccessForwardingRuleActionbypass,
		"forward": NetworkaccessForwardingRuleActionforward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingRuleAction(input)
	return &out, nil
}
