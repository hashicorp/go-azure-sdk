package availabilitygrouplisteners

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Commit string

const (
	CommitAsynchronousCommit Commit = "Asynchronous_Commit"
	CommitSynchronousCommit  Commit = "Synchronous_Commit"
)

func PossibleValuesForCommit() []string {
	return []string{
		string(CommitAsynchronousCommit),
		string(CommitSynchronousCommit),
	}
}

type Failover string

const (
	FailoverAutomatic Failover = "Automatic"
	FailoverManual    Failover = "Manual"
)

func PossibleValuesForFailover() []string {
	return []string{
		string(FailoverAutomatic),
		string(FailoverManual),
	}
}

type ReadableSecondary string

const (
	ReadableSecondaryAll      ReadableSecondary = "All"
	ReadableSecondaryNo       ReadableSecondary = "No"
	ReadableSecondaryReadOnly ReadableSecondary = "Read_Only"
)

func PossibleValuesForReadableSecondary() []string {
	return []string{
		string(ReadableSecondaryAll),
		string(ReadableSecondaryNo),
		string(ReadableSecondaryReadOnly),
	}
}

type Role string

const (
	RolePrimary   Role = "Primary"
	RoleSecondary Role = "Secondary"
)

func PossibleValuesForRole() []string {
	return []string{
		string(RolePrimary),
		string(RoleSecondary),
	}
}
