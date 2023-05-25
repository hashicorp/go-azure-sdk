package distributedavailabilitygroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstanceRole string

const (
	InstanceRolePrimary   InstanceRole = "Primary"
	InstanceRoleSecondary InstanceRole = "Secondary"
)

func PossibleValuesForInstanceRole() []string {
	return []string{
		string(InstanceRolePrimary),
		string(InstanceRoleSecondary),
	}
}

func (s *InstanceRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstanceRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstanceRole(input string) (*InstanceRole, error) {
	vals := map[string]InstanceRole{
		"primary":   InstanceRolePrimary,
		"secondary": InstanceRoleSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstanceRole(input)
	return &out, nil
}

type ReplicationMode string

const (
	ReplicationModeAsync ReplicationMode = "Async"
	ReplicationModeSync  ReplicationMode = "Sync"
)

func PossibleValuesForReplicationMode() []string {
	return []string{
		string(ReplicationModeAsync),
		string(ReplicationModeSync),
	}
}

func (s *ReplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicationMode(input string) (*ReplicationMode, error) {
	vals := map[string]ReplicationMode{
		"async": ReplicationModeAsync,
		"sync":  ReplicationModeSync,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationMode(input)
	return &out, nil
}

type RoleChangeType string

const (
	RoleChangeTypeForced  RoleChangeType = "Forced"
	RoleChangeTypePlanned RoleChangeType = "Planned"
)

func PossibleValuesForRoleChangeType() []string {
	return []string{
		string(RoleChangeTypeForced),
		string(RoleChangeTypePlanned),
	}
}

func (s *RoleChangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleChangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleChangeType(input string) (*RoleChangeType, error) {
	vals := map[string]RoleChangeType{
		"forced":  RoleChangeTypeForced,
		"planned": RoleChangeTypePlanned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleChangeType(input)
	return &out, nil
}
