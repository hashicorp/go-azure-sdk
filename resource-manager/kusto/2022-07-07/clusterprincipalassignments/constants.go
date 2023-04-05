package clusterprincipalassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterPrincipalRole string

const (
	ClusterPrincipalRoleAllDatabasesAdmin  ClusterPrincipalRole = "AllDatabasesAdmin"
	ClusterPrincipalRoleAllDatabasesViewer ClusterPrincipalRole = "AllDatabasesViewer"
)

func PossibleValuesForClusterPrincipalRole() []string {
	return []string{
		string(ClusterPrincipalRoleAllDatabasesAdmin),
		string(ClusterPrincipalRoleAllDatabasesViewer),
	}
}

type PrincipalAssignmentType string

const (
	PrincipalAssignmentTypeMicrosoftPointKustoClustersPrincipalAssignments PrincipalAssignmentType = "Microsoft.Kusto/clusters/principalAssignments"
)

func PossibleValuesForPrincipalAssignmentType() []string {
	return []string{
		string(PrincipalAssignmentTypeMicrosoftPointKustoClustersPrincipalAssignments),
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
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
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
