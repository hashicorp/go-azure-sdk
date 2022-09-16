package sqlpoolsreplicationlinks

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationRole string

const (
	ReplicationRoleCopy                 ReplicationRole = "Copy"
	ReplicationRoleNonReadableSecondary ReplicationRole = "NonReadableSecondary"
	ReplicationRolePrimary              ReplicationRole = "Primary"
	ReplicationRoleSecondary            ReplicationRole = "Secondary"
	ReplicationRoleSource               ReplicationRole = "Source"
)

func PossibleValuesForReplicationRole() []string {
	return []string{
		string(ReplicationRoleCopy),
		string(ReplicationRoleNonReadableSecondary),
		string(ReplicationRolePrimary),
		string(ReplicationRoleSecondary),
		string(ReplicationRoleSource),
	}
}

func parseReplicationRole(input string) (*ReplicationRole, error) {
	vals := map[string]ReplicationRole{
		"copy":                 ReplicationRoleCopy,
		"nonreadablesecondary": ReplicationRoleNonReadableSecondary,
		"primary":              ReplicationRolePrimary,
		"secondary":            ReplicationRoleSecondary,
		"source":               ReplicationRoleSource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationRole(input)
	return &out, nil
}

type ReplicationState string

const (
	ReplicationStateCATCHUP   ReplicationState = "CATCH_UP"
	ReplicationStatePENDING   ReplicationState = "PENDING"
	ReplicationStateSEEDING   ReplicationState = "SEEDING"
	ReplicationStateSUSPENDED ReplicationState = "SUSPENDED"
)

func PossibleValuesForReplicationState() []string {
	return []string{
		string(ReplicationStateCATCHUP),
		string(ReplicationStatePENDING),
		string(ReplicationStateSEEDING),
		string(ReplicationStateSUSPENDED),
	}
}

func parseReplicationState(input string) (*ReplicationState, error) {
	vals := map[string]ReplicationState{
		"catch_up":  ReplicationStateCATCHUP,
		"pending":   ReplicationStatePENDING,
		"seeding":   ReplicationStateSEEDING,
		"suspended": ReplicationStateSUSPENDED,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationState(input)
	return &out, nil
}
