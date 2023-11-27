package failovergroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailoverGroupReplicationRole string

const (
	FailoverGroupReplicationRolePrimary   FailoverGroupReplicationRole = "Primary"
	FailoverGroupReplicationRoleSecondary FailoverGroupReplicationRole = "Secondary"
)

func PossibleValuesForFailoverGroupReplicationRole() []string {
	return []string{
		string(FailoverGroupReplicationRolePrimary),
		string(FailoverGroupReplicationRoleSecondary),
	}
}

func (s *FailoverGroupReplicationRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFailoverGroupReplicationRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFailoverGroupReplicationRole(input string) (*FailoverGroupReplicationRole, error) {
	vals := map[string]FailoverGroupReplicationRole{
		"primary":   FailoverGroupReplicationRolePrimary,
		"secondary": FailoverGroupReplicationRoleSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FailoverGroupReplicationRole(input)
	return &out, nil
}

type ReadOnlyEndpointFailoverPolicy string

const (
	ReadOnlyEndpointFailoverPolicyDisabled ReadOnlyEndpointFailoverPolicy = "Disabled"
	ReadOnlyEndpointFailoverPolicyEnabled  ReadOnlyEndpointFailoverPolicy = "Enabled"
)

func PossibleValuesForReadOnlyEndpointFailoverPolicy() []string {
	return []string{
		string(ReadOnlyEndpointFailoverPolicyDisabled),
		string(ReadOnlyEndpointFailoverPolicyEnabled),
	}
}

func (s *ReadOnlyEndpointFailoverPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReadOnlyEndpointFailoverPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReadOnlyEndpointFailoverPolicy(input string) (*ReadOnlyEndpointFailoverPolicy, error) {
	vals := map[string]ReadOnlyEndpointFailoverPolicy{
		"disabled": ReadOnlyEndpointFailoverPolicyDisabled,
		"enabled":  ReadOnlyEndpointFailoverPolicyEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReadOnlyEndpointFailoverPolicy(input)
	return &out, nil
}

type ReadWriteEndpointFailoverPolicy string

const (
	ReadWriteEndpointFailoverPolicyAutomatic ReadWriteEndpointFailoverPolicy = "Automatic"
	ReadWriteEndpointFailoverPolicyManual    ReadWriteEndpointFailoverPolicy = "Manual"
)

func PossibleValuesForReadWriteEndpointFailoverPolicy() []string {
	return []string{
		string(ReadWriteEndpointFailoverPolicyAutomatic),
		string(ReadWriteEndpointFailoverPolicyManual),
	}
}

func (s *ReadWriteEndpointFailoverPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReadWriteEndpointFailoverPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReadWriteEndpointFailoverPolicy(input string) (*ReadWriteEndpointFailoverPolicy, error) {
	vals := map[string]ReadWriteEndpointFailoverPolicy{
		"automatic": ReadWriteEndpointFailoverPolicyAutomatic,
		"manual":    ReadWriteEndpointFailoverPolicyManual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReadWriteEndpointFailoverPolicy(input)
	return &out, nil
}
