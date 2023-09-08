package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRouteRoutingType string

const (
	CallRouteRoutingTypeforwarded CallRouteRoutingType = "Forwarded"
	CallRouteRoutingTypelookup    CallRouteRoutingType = "Lookup"
	CallRouteRoutingTypeselfFork  CallRouteRoutingType = "SelfFork"
)

func PossibleValuesForCallRouteRoutingType() []string {
	return []string{
		string(CallRouteRoutingTypeforwarded),
		string(CallRouteRoutingTypelookup),
		string(CallRouteRoutingTypeselfFork),
	}
}

func parseCallRouteRoutingType(input string) (*CallRouteRoutingType, error) {
	vals := map[string]CallRouteRoutingType{
		"forwarded": CallRouteRoutingTypeforwarded,
		"lookup":    CallRouteRoutingTypelookup,
		"selffork":  CallRouteRoutingTypeselfFork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRouteRoutingType(input)
	return &out, nil
}
