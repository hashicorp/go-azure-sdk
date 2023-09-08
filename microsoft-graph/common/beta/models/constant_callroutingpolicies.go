package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRoutingPolicies string

const (
	CallRoutingPoliciesdisableForwarding            CallRoutingPolicies = "DisableForwarding"
	CallRoutingPoliciesdisableForwardingExceptPhone CallRoutingPolicies = "DisableForwardingExceptPhone"
	CallRoutingPoliciesnoMissedCall                 CallRoutingPolicies = "NoMissedCall"
	CallRoutingPoliciesnone                         CallRoutingPolicies = "None"
	CallRoutingPoliciespreferSkypeForBusiness       CallRoutingPolicies = "PreferSkypeForBusiness"
)

func PossibleValuesForCallRoutingPolicies() []string {
	return []string{
		string(CallRoutingPoliciesdisableForwarding),
		string(CallRoutingPoliciesdisableForwardingExceptPhone),
		string(CallRoutingPoliciesnoMissedCall),
		string(CallRoutingPoliciesnone),
		string(CallRoutingPoliciespreferSkypeForBusiness),
	}
}

func parseCallRoutingPolicies(input string) (*CallRoutingPolicies, error) {
	vals := map[string]CallRoutingPolicies{
		"disableforwarding":            CallRoutingPoliciesdisableForwarding,
		"disableforwardingexceptphone": CallRoutingPoliciesdisableForwardingExceptPhone,
		"nomissedcall":                 CallRoutingPoliciesnoMissedCall,
		"none":                         CallRoutingPoliciesnone,
		"preferskypeforbusiness":       CallRoutingPoliciespreferSkypeForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRoutingPolicies(input)
	return &out, nil
}
