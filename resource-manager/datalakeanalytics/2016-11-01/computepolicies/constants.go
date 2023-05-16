package computepolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AADObjectType string

const (
	AADObjectTypeGroup            AADObjectType = "Group"
	AADObjectTypeServicePrincipal AADObjectType = "ServicePrincipal"
	AADObjectTypeUser             AADObjectType = "User"
)

func PossibleValuesForAADObjectType() []string {
	return []string{
		string(AADObjectTypeGroup),
		string(AADObjectTypeServicePrincipal),
		string(AADObjectTypeUser),
	}
}

func (s *AADObjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAADObjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAADObjectType(input string) (*AADObjectType, error) {
	vals := map[string]AADObjectType{
		"group":            AADObjectTypeGroup,
		"serviceprincipal": AADObjectTypeServicePrincipal,
		"user":             AADObjectTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AADObjectType(input)
	return &out, nil
}
