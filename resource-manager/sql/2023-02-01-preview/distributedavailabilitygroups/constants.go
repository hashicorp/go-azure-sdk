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

type LinkRole string

const (
	LinkRolePrimary   LinkRole = "Primary"
	LinkRoleSecondary LinkRole = "Secondary"
)

func PossibleValuesForLinkRole() []string {
	return []string{
		string(LinkRolePrimary),
		string(LinkRoleSecondary),
	}
}

func (s *LinkRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLinkRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLinkRole(input string) (*LinkRole, error) {
	vals := map[string]LinkRole{
		"primary":   LinkRolePrimary,
		"secondary": LinkRoleSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LinkRole(input)
	return &out, nil
}

type ReplicaConnectedState string

const (
	ReplicaConnectedStateCONNECTED    ReplicaConnectedState = "CONNECTED"
	ReplicaConnectedStateDISCONNECTED ReplicaConnectedState = "DISCONNECTED"
)

func PossibleValuesForReplicaConnectedState() []string {
	return []string{
		string(ReplicaConnectedStateCONNECTED),
		string(ReplicaConnectedStateDISCONNECTED),
	}
}

func (s *ReplicaConnectedState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicaConnectedState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicaConnectedState(input string) (*ReplicaConnectedState, error) {
	vals := map[string]ReplicaConnectedState{
		"connected":    ReplicaConnectedStateCONNECTED,
		"disconnected": ReplicaConnectedStateDISCONNECTED,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicaConnectedState(input)
	return &out, nil
}

type ReplicaSynchronizationHealth string

const (
	ReplicaSynchronizationHealthHEALTHY          ReplicaSynchronizationHealth = "HEALTHY"
	ReplicaSynchronizationHealthNOTHEALTHY       ReplicaSynchronizationHealth = "NOT_HEALTHY"
	ReplicaSynchronizationHealthPARTIALLYHEALTHY ReplicaSynchronizationHealth = "PARTIALLY_HEALTHY"
)

func PossibleValuesForReplicaSynchronizationHealth() []string {
	return []string{
		string(ReplicaSynchronizationHealthHEALTHY),
		string(ReplicaSynchronizationHealthNOTHEALTHY),
		string(ReplicaSynchronizationHealthPARTIALLYHEALTHY),
	}
}

func (s *ReplicaSynchronizationHealth) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicaSynchronizationHealth(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicaSynchronizationHealth(input string) (*ReplicaSynchronizationHealth, error) {
	vals := map[string]ReplicaSynchronizationHealth{
		"healthy":           ReplicaSynchronizationHealthHEALTHY,
		"not_healthy":       ReplicaSynchronizationHealthNOTHEALTHY,
		"partially_healthy": ReplicaSynchronizationHealthPARTIALLYHEALTHY,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicaSynchronizationHealth(input)
	return &out, nil
}

type ReplicationModeType string

const (
	ReplicationModeTypeAsync ReplicationModeType = "Async"
	ReplicationModeTypeSync  ReplicationModeType = "Sync"
)

func PossibleValuesForReplicationModeType() []string {
	return []string{
		string(ReplicationModeTypeAsync),
		string(ReplicationModeTypeSync),
	}
}

func (s *ReplicationModeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicationModeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicationModeType(input string) (*ReplicationModeType, error) {
	vals := map[string]ReplicationModeType{
		"async": ReplicationModeTypeAsync,
		"sync":  ReplicationModeTypeSync,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationModeType(input)
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
