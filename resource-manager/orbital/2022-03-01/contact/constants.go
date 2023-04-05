package contact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactsStatus string

const (
	ContactsStatusCancelled         ContactsStatus = "cancelled"
	ContactsStatusFailed            ContactsStatus = "failed"
	ContactsStatusProviderCancelled ContactsStatus = "providerCancelled"
	ContactsStatusScheduled         ContactsStatus = "scheduled"
	ContactsStatusSucceeded         ContactsStatus = "succeeded"
)

func PossibleValuesForContactsStatus() []string {
	return []string{
		string(ContactsStatusCancelled),
		string(ContactsStatusFailed),
		string(ContactsStatusProviderCancelled),
		string(ContactsStatusScheduled),
		string(ContactsStatusSucceeded),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}
