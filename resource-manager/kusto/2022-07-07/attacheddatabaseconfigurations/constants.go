package attacheddatabaseconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachedDatabaseType string

const (
	AttachedDatabaseTypeMicrosoftPointKustoClustersAttachedDatabaseConfigurations AttachedDatabaseType = "Microsoft.Kusto/clusters/attachedDatabaseConfigurations"
)

func PossibleValuesForAttachedDatabaseType() []string {
	return []string{
		string(AttachedDatabaseTypeMicrosoftPointKustoClustersAttachedDatabaseConfigurations),
	}
}

type DefaultPrincipalsModificationKind string

const (
	DefaultPrincipalsModificationKindNone    DefaultPrincipalsModificationKind = "None"
	DefaultPrincipalsModificationKindReplace DefaultPrincipalsModificationKind = "Replace"
	DefaultPrincipalsModificationKindUnion   DefaultPrincipalsModificationKind = "Union"
)

func PossibleValuesForDefaultPrincipalsModificationKind() []string {
	return []string{
		string(DefaultPrincipalsModificationKindNone),
		string(DefaultPrincipalsModificationKindReplace),
		string(DefaultPrincipalsModificationKindUnion),
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
