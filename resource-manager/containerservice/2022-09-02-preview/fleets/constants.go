package fleets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FleetProvisioningState string

const (
	FleetProvisioningStateCanceled  FleetProvisioningState = "Canceled"
	FleetProvisioningStateCreating  FleetProvisioningState = "Creating"
	FleetProvisioningStateDeleting  FleetProvisioningState = "Deleting"
	FleetProvisioningStateFailed    FleetProvisioningState = "Failed"
	FleetProvisioningStateSucceeded FleetProvisioningState = "Succeeded"
	FleetProvisioningStateUpdating  FleetProvisioningState = "Updating"
)

func PossibleValuesForFleetProvisioningState() []string {
	return []string{
		string(FleetProvisioningStateCanceled),
		string(FleetProvisioningStateCreating),
		string(FleetProvisioningStateDeleting),
		string(FleetProvisioningStateFailed),
		string(FleetProvisioningStateSucceeded),
		string(FleetProvisioningStateUpdating),
	}
}
