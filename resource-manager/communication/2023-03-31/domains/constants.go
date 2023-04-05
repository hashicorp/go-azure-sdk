package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainManagement string

const (
	DomainManagementAzureManaged                    DomainManagement = "AzureManaged"
	DomainManagementCustomerManaged                 DomainManagement = "CustomerManaged"
	DomainManagementCustomerManagedInExchangeOnline DomainManagement = "CustomerManagedInExchangeOnline"
)

func PossibleValuesForDomainManagement() []string {
	return []string{
		string(DomainManagementAzureManaged),
		string(DomainManagementCustomerManaged),
		string(DomainManagementCustomerManagedInExchangeOnline),
	}
}

type DomainsProvisioningState string

const (
	DomainsProvisioningStateCanceled  DomainsProvisioningState = "Canceled"
	DomainsProvisioningStateCreating  DomainsProvisioningState = "Creating"
	DomainsProvisioningStateDeleting  DomainsProvisioningState = "Deleting"
	DomainsProvisioningStateFailed    DomainsProvisioningState = "Failed"
	DomainsProvisioningStateMoving    DomainsProvisioningState = "Moving"
	DomainsProvisioningStateRunning   DomainsProvisioningState = "Running"
	DomainsProvisioningStateSucceeded DomainsProvisioningState = "Succeeded"
	DomainsProvisioningStateUnknown   DomainsProvisioningState = "Unknown"
	DomainsProvisioningStateUpdating  DomainsProvisioningState = "Updating"
)

func PossibleValuesForDomainsProvisioningState() []string {
	return []string{
		string(DomainsProvisioningStateCanceled),
		string(DomainsProvisioningStateCreating),
		string(DomainsProvisioningStateDeleting),
		string(DomainsProvisioningStateFailed),
		string(DomainsProvisioningStateMoving),
		string(DomainsProvisioningStateRunning),
		string(DomainsProvisioningStateSucceeded),
		string(DomainsProvisioningStateUnknown),
		string(DomainsProvisioningStateUpdating),
	}
}

type UserEngagementTracking string

const (
	UserEngagementTrackingDisabled UserEngagementTracking = "Disabled"
	UserEngagementTrackingEnabled  UserEngagementTracking = "Enabled"
)

func PossibleValuesForUserEngagementTracking() []string {
	return []string{
		string(UserEngagementTrackingDisabled),
		string(UserEngagementTrackingEnabled),
	}
}

type VerificationStatus string

const (
	VerificationStatusCancellationRequested  VerificationStatus = "CancellationRequested"
	VerificationStatusNotStarted             VerificationStatus = "NotStarted"
	VerificationStatusVerificationFailed     VerificationStatus = "VerificationFailed"
	VerificationStatusVerificationInProgress VerificationStatus = "VerificationInProgress"
	VerificationStatusVerificationRequested  VerificationStatus = "VerificationRequested"
	VerificationStatusVerified               VerificationStatus = "Verified"
)

func PossibleValuesForVerificationStatus() []string {
	return []string{
		string(VerificationStatusCancellationRequested),
		string(VerificationStatusNotStarted),
		string(VerificationStatusVerificationFailed),
		string(VerificationStatusVerificationInProgress),
		string(VerificationStatusVerificationRequested),
		string(VerificationStatusVerified),
	}
}

type VerificationType string

const (
	VerificationTypeDKIM    VerificationType = "DKIM"
	VerificationTypeDKIMTwo VerificationType = "DKIM2"
	VerificationTypeDMARC   VerificationType = "DMARC"
	VerificationTypeDomain  VerificationType = "Domain"
	VerificationTypeSPF     VerificationType = "SPF"
)

func PossibleValuesForVerificationType() []string {
	return []string{
		string(VerificationTypeDKIM),
		string(VerificationTypeDKIMTwo),
		string(VerificationTypeDMARC),
		string(VerificationTypeDomain),
		string(VerificationTypeSPF),
	}
}
