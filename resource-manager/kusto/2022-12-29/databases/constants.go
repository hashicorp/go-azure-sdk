package databases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallerRole string

const (
	CallerRoleAdmin CallerRole = "Admin"
	CallerRoleNone  CallerRole = "None"
)

func PossibleValuesForCallerRole() []string {
	return []string{
		string(CallerRoleAdmin),
		string(CallerRoleNone),
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

type DatabasePrincipalType string

const (
	DatabasePrincipalTypeApp   DatabasePrincipalType = "App"
	DatabasePrincipalTypeGroup DatabasePrincipalType = "Group"
	DatabasePrincipalTypeUser  DatabasePrincipalType = "User"
)

func PossibleValuesForDatabasePrincipalType() []string {
	return []string{
		string(DatabasePrincipalTypeApp),
		string(DatabasePrincipalTypeGroup),
		string(DatabasePrincipalTypeUser),
	}
}

type DatabaseShareOrigin string

const (
	DatabaseShareOriginDataShare DatabaseShareOrigin = "DataShare"
	DatabaseShareOriginDirect    DatabaseShareOrigin = "Direct"
	DatabaseShareOriginOther     DatabaseShareOrigin = "Other"
)

func PossibleValuesForDatabaseShareOrigin() []string {
	return []string{
		string(DatabaseShareOriginDataShare),
		string(DatabaseShareOriginDirect),
		string(DatabaseShareOriginOther),
	}
}

type Kind string

const (
	KindReadOnlyFollowing Kind = "ReadOnlyFollowing"
	KindReadWrite         Kind = "ReadWrite"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindReadOnlyFollowing),
		string(KindReadWrite),
	}
}

type PrincipalsModificationKind string

const (
	PrincipalsModificationKindNone    PrincipalsModificationKind = "None"
	PrincipalsModificationKindReplace PrincipalsModificationKind = "Replace"
	PrincipalsModificationKindUnion   PrincipalsModificationKind = "Union"
)

func PossibleValuesForPrincipalsModificationKind() []string {
	return []string{
		string(PrincipalsModificationKindNone),
		string(PrincipalsModificationKindReplace),
		string(PrincipalsModificationKindUnion),
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

type Type string

const (
	TypeMicrosoftPointKustoClustersAttachedDatabaseConfigurations Type = "Microsoft.Kusto/clusters/attachedDatabaseConfigurations"
	TypeMicrosoftPointKustoClustersDatabases                      Type = "Microsoft.Kusto/clusters/databases"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointKustoClustersAttachedDatabaseConfigurations),
		string(TypeMicrosoftPointKustoClustersDatabases),
	}
}
