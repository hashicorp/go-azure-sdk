package networkfeaturesslotoperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RouteType string

const (
	RouteTypeDefault   RouteType = "DEFAULT"
	RouteTypeINHERITED RouteType = "INHERITED"
	RouteTypeSTATIC    RouteType = "STATIC"
)

func PossibleValuesForRouteType() []string {
	return []string{
		string(RouteTypeDefault),
		string(RouteTypeINHERITED),
		string(RouteTypeSTATIC),
	}
}

func (s *RouteType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRouteType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRouteType(input string) (*RouteType, error) {
	vals := map[string]RouteType{
		"default":   RouteTypeDefault,
		"inherited": RouteTypeINHERITED,
		"static":    RouteTypeSTATIC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RouteType(input)
	return &out, nil
}
