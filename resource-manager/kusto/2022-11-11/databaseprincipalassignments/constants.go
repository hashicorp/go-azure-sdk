package databaseprincipalassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabasePrincipalAssignmentType string

const (
	DatabasePrincipalAssignmentTypeMicrosoftPointKustoClustersDatabasesPrincipalAssignments DatabasePrincipalAssignmentType = "Microsoft.Kusto/clusters/databases/principalAssignments"
)

func PossibleValuesForDatabasePrincipalAssignmentType() []string {
	return []string{
		string(DatabasePrincipalAssignmentTypeMicrosoftPointKustoClustersDatabasesPrincipalAssignments),
	}
}

type DatabasePrincipalRole string

const (
	DatabasePrincipalRoleAdmin              DatabasePrincipalRole = "Admin"
	DatabasePrincipalRoleIngestor           DatabasePrincipalRole = "Ingestor"
	DatabasePrincipalRoleMonitor            DatabasePrincipalRole = "Monitor"
	DatabasePrincipalRoleUnrestrictedViewer DatabasePrincipalRole = "UnrestrictedViewer"
	DatabasePrincipalRoleUser               DatabasePrincipalRole = "User"
	DatabasePrincipalRoleViewer             DatabasePrincipalRole = "Viewer"
)

func PossibleValuesForDatabasePrincipalRole() []string {
	return []string{
		string(DatabasePrincipalRoleAdmin),
		string(DatabasePrincipalRoleIngestor),
		string(DatabasePrincipalRoleMonitor),
		string(DatabasePrincipalRoleUnrestrictedViewer),
		string(DatabasePrincipalRoleUser),
		string(DatabasePrincipalRoleViewer),
	}
}

type PrincipalType string

const (
	PrincipalTypeApp   PrincipalType = "App"
	PrincipalTypeGroup PrincipalType = "Group"
	PrincipalTypeUser  PrincipalType = "User"
)

func PossibleValuesForPrincipalType() []string {
	return []string{
		string(PrincipalTypeApp),
		string(PrincipalTypeGroup),
		string(PrincipalTypeUser),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
	}
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
	}
}
