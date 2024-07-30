package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRouteRoutingType string

const (
	CallRouteRoutingType_Forwarded CallRouteRoutingType = "forwarded"
	CallRouteRoutingType_Lookup    CallRouteRoutingType = "lookup"
	CallRouteRoutingType_SelfFork  CallRouteRoutingType = "selfFork"
)

func PossibleValuesForCallRouteRoutingType() []string {
	return []string{
		string(CallRouteRoutingType_Forwarded),
		string(CallRouteRoutingType_Lookup),
		string(CallRouteRoutingType_SelfFork),
	}
}

func (s *CallRouteRoutingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRouteRoutingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRouteRoutingType(input string) (*CallRouteRoutingType, error) {
	vals := map[string]CallRouteRoutingType{
		"forwarded": CallRouteRoutingType_Forwarded,
		"lookup":    CallRouteRoutingType_Lookup,
		"selffork":  CallRouteRoutingType_SelfFork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRouteRoutingType(input)
	return &out, nil
}
