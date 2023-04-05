package communicationservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

type CommunicationServicesProvisioningState string

const (
	CommunicationServicesProvisioningStateCanceled  CommunicationServicesProvisioningState = "Canceled"
	CommunicationServicesProvisioningStateCreating  CommunicationServicesProvisioningState = "Creating"
	CommunicationServicesProvisioningStateDeleting  CommunicationServicesProvisioningState = "Deleting"
	CommunicationServicesProvisioningStateFailed    CommunicationServicesProvisioningState = "Failed"
	CommunicationServicesProvisioningStateMoving    CommunicationServicesProvisioningState = "Moving"
	CommunicationServicesProvisioningStateRunning   CommunicationServicesProvisioningState = "Running"
	CommunicationServicesProvisioningStateSucceeded CommunicationServicesProvisioningState = "Succeeded"
	CommunicationServicesProvisioningStateUnknown   CommunicationServicesProvisioningState = "Unknown"
	CommunicationServicesProvisioningStateUpdating  CommunicationServicesProvisioningState = "Updating"
)

func PossibleValuesForCommunicationServicesProvisioningState() []string {
	return []string{
		string(CommunicationServicesProvisioningStateCanceled),
		string(CommunicationServicesProvisioningStateCreating),
		string(CommunicationServicesProvisioningStateDeleting),
		string(CommunicationServicesProvisioningStateFailed),
		string(CommunicationServicesProvisioningStateMoving),
		string(CommunicationServicesProvisioningStateRunning),
		string(CommunicationServicesProvisioningStateSucceeded),
		string(CommunicationServicesProvisioningStateUnknown),
		string(CommunicationServicesProvisioningStateUpdating),
	}
}

type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSecondary KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}
