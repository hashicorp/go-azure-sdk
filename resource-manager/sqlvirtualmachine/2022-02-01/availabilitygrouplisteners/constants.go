package availabilitygrouplisteners

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Commit string

const (
	CommitASYNCHRONOUSCOMMIT Commit = "ASYNCHRONOUS_COMMIT"
	CommitSYNCHRONOUSCOMMIT  Commit = "SYNCHRONOUS_COMMIT"
)

func PossibleValuesForCommit() []string {
	return []string{
		string(CommitASYNCHRONOUSCOMMIT),
		string(CommitSYNCHRONOUSCOMMIT),
	}
}

func parseCommit(input string) (*Commit, error) {
	vals := map[string]Commit{
		"asynchronous_commit": CommitASYNCHRONOUSCOMMIT,
		"synchronous_commit":  CommitSYNCHRONOUSCOMMIT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Commit(input)
	return &out, nil
}

type Failover string

const (
	FailoverAUTOMATIC Failover = "AUTOMATIC"
	FailoverMANUAL    Failover = "MANUAL"
)

func PossibleValuesForFailover() []string {
	return []string{
		string(FailoverAUTOMATIC),
		string(FailoverMANUAL),
	}
}

func parseFailover(input string) (*Failover, error) {
	vals := map[string]Failover{
		"automatic": FailoverAUTOMATIC,
		"manual":    FailoverMANUAL,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Failover(input)
	return &out, nil
}

type ReadableSecondary string

const (
	ReadableSecondaryALL      ReadableSecondary = "ALL"
	ReadableSecondaryNO       ReadableSecondary = "NO"
	ReadableSecondaryREADONLY ReadableSecondary = "READ_ONLY"
)

func PossibleValuesForReadableSecondary() []string {
	return []string{
		string(ReadableSecondaryALL),
		string(ReadableSecondaryNO),
		string(ReadableSecondaryREADONLY),
	}
}

func parseReadableSecondary(input string) (*ReadableSecondary, error) {
	vals := map[string]ReadableSecondary{
		"all":       ReadableSecondaryALL,
		"no":        ReadableSecondaryNO,
		"read_only": ReadableSecondaryREADONLY,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReadableSecondary(input)
	return &out, nil
}

type Role string

const (
	RolePRIMARY   Role = "PRIMARY"
	RoleSECONDARY Role = "SECONDARY"
)

func PossibleValuesForRole() []string {
	return []string{
		string(RolePRIMARY),
		string(RoleSECONDARY),
	}
}

func parseRole(input string) (*Role, error) {
	vals := map[string]Role{
		"primary":   RolePRIMARY,
		"secondary": RoleSECONDARY,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Role(input)
	return &out, nil
}
