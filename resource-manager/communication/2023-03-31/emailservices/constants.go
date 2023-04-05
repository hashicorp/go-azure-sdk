package emailservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailServicesProvisioningState string

const (
	EmailServicesProvisioningStateCanceled  EmailServicesProvisioningState = "Canceled"
	EmailServicesProvisioningStateCreating  EmailServicesProvisioningState = "Creating"
	EmailServicesProvisioningStateDeleting  EmailServicesProvisioningState = "Deleting"
	EmailServicesProvisioningStateFailed    EmailServicesProvisioningState = "Failed"
	EmailServicesProvisioningStateMoving    EmailServicesProvisioningState = "Moving"
	EmailServicesProvisioningStateRunning   EmailServicesProvisioningState = "Running"
	EmailServicesProvisioningStateSucceeded EmailServicesProvisioningState = "Succeeded"
	EmailServicesProvisioningStateUnknown   EmailServicesProvisioningState = "Unknown"
	EmailServicesProvisioningStateUpdating  EmailServicesProvisioningState = "Updating"
)

func PossibleValuesForEmailServicesProvisioningState() []string {
	return []string{
		string(EmailServicesProvisioningStateCanceled),
		string(EmailServicesProvisioningStateCreating),
		string(EmailServicesProvisioningStateDeleting),
		string(EmailServicesProvisioningStateFailed),
		string(EmailServicesProvisioningStateMoving),
		string(EmailServicesProvisioningStateRunning),
		string(EmailServicesProvisioningStateSucceeded),
		string(EmailServicesProvisioningStateUnknown),
		string(EmailServicesProvisioningStateUpdating),
	}
}
