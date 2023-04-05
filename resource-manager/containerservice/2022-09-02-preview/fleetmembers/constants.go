package fleetmembers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FleetMemberProvisioningState string

const (
	FleetMemberProvisioningStateCanceled  FleetMemberProvisioningState = "Canceled"
	FleetMemberProvisioningStateFailed    FleetMemberProvisioningState = "Failed"
	FleetMemberProvisioningStateJoining   FleetMemberProvisioningState = "Joining"
	FleetMemberProvisioningStateLeaving   FleetMemberProvisioningState = "Leaving"
	FleetMemberProvisioningStateSucceeded FleetMemberProvisioningState = "Succeeded"
	FleetMemberProvisioningStateUpdating  FleetMemberProvisioningState = "Updating"
)

func PossibleValuesForFleetMemberProvisioningState() []string {
	return []string{
		string(FleetMemberProvisioningStateCanceled),
		string(FleetMemberProvisioningStateFailed),
		string(FleetMemberProvisioningStateJoining),
		string(FleetMemberProvisioningStateLeaving),
		string(FleetMemberProvisioningStateSucceeded),
		string(FleetMemberProvisioningStateUpdating),
	}
}
