package roleassignments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrincipalType string

const (
	PrincipalTypeForeignGroup     PrincipalType = "ForeignGroup"
	PrincipalTypeGroup            PrincipalType = "Group"
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	PrincipalTypeUser             PrincipalType = "User"
)

func PossibleValuesForPrincipalType() []string {
	return []string{
		string(PrincipalTypeForeignGroup),
		string(PrincipalTypeGroup),
		string(PrincipalTypeServicePrincipal),
		string(PrincipalTypeUser),
	}
}

func (s *PrincipalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrincipalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrincipalType(input string) (*PrincipalType, error) {
	vals := map[string]PrincipalType{
		"foreigngroup":     PrincipalTypeForeignGroup,
		"group":            PrincipalTypeGroup,
		"serviceprincipal": PrincipalTypeServicePrincipal,
		"user":             PrincipalTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrincipalType(input)
	return &out, nil
}
