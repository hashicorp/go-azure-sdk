package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyRuleDeltaAction string

const (
	NetworkaccessPolicyRuleDeltaActionbypass  NetworkaccessPolicyRuleDeltaAction = "Bypass"
	NetworkaccessPolicyRuleDeltaActionforward NetworkaccessPolicyRuleDeltaAction = "Forward"
)

func PossibleValuesForNetworkaccessPolicyRuleDeltaAction() []string {
	return []string{
		string(NetworkaccessPolicyRuleDeltaActionbypass),
		string(NetworkaccessPolicyRuleDeltaActionforward),
	}
}

func parseNetworkaccessPolicyRuleDeltaAction(input string) (*NetworkaccessPolicyRuleDeltaAction, error) {
	vals := map[string]NetworkaccessPolicyRuleDeltaAction{
		"bypass":  NetworkaccessPolicyRuleDeltaActionbypass,
		"forward": NetworkaccessPolicyRuleDeltaActionforward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPolicyRuleDeltaAction(input)
	return &out, nil
}
