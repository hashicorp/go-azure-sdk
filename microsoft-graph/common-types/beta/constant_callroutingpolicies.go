package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRoutingPolicies string

const (
	CallRoutingPolicies_DisableForwarding            CallRoutingPolicies = "disableForwarding"
	CallRoutingPolicies_DisableForwardingExceptPhone CallRoutingPolicies = "disableForwardingExceptPhone"
	CallRoutingPolicies_NoMissedCall                 CallRoutingPolicies = "noMissedCall"
	CallRoutingPolicies_None                         CallRoutingPolicies = "none"
	CallRoutingPolicies_PreferSkypeForBusiness       CallRoutingPolicies = "preferSkypeForBusiness"
)

func PossibleValuesForCallRoutingPolicies() []string {
	return []string{
		string(CallRoutingPolicies_DisableForwarding),
		string(CallRoutingPolicies_DisableForwardingExceptPhone),
		string(CallRoutingPolicies_NoMissedCall),
		string(CallRoutingPolicies_None),
		string(CallRoutingPolicies_PreferSkypeForBusiness),
	}
}

func (s *CallRoutingPolicies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRoutingPolicies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRoutingPolicies(input string) (*CallRoutingPolicies, error) {
	vals := map[string]CallRoutingPolicies{
		"disableforwarding":            CallRoutingPolicies_DisableForwarding,
		"disableforwardingexceptphone": CallRoutingPolicies_DisableForwardingExceptPhone,
		"nomissedcall":                 CallRoutingPolicies_NoMissedCall,
		"none":                         CallRoutingPolicies_None,
		"preferskypeforbusiness":       CallRoutingPolicies_PreferSkypeForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRoutingPolicies(input)
	return &out, nil
}
