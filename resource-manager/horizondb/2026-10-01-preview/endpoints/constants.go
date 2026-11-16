package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointRole string

const (
	EndpointRoleReadOnly  EndpointRole = "ReadOnly"
	EndpointRoleReadWrite EndpointRole = "ReadWrite"
)

func PossibleValuesForEndpointRole() []string {
	return []string{
		string(EndpointRoleReadOnly),
		string(EndpointRoleReadWrite),
	}
}

func (s *EndpointRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointRole(input string) (*EndpointRole, error) {
	vals := map[string]EndpointRole{
		"readonly":  EndpointRoleReadOnly,
		"readwrite": EndpointRoleReadWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointRole(input)
	return &out, nil
}

type EndpointScope string

const (
	EndpointScopeCluster EndpointScope = "Cluster"
	EndpointScopePool    EndpointScope = "Pool"
)

func PossibleValuesForEndpointScope() []string {
	return []string{
		string(EndpointScopeCluster),
		string(EndpointScopePool),
	}
}

func (s *EndpointScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointScope(input string) (*EndpointScope, error) {
	vals := map[string]EndpointScope{
		"cluster": EndpointScopeCluster,
		"pool":    EndpointScopePool,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointScope(input)
	return &out, nil
}
