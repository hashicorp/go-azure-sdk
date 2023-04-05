package spacecraft

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Direction string

const (
	DirectionDownlink Direction = "downlink"
	DirectionUplink   Direction = "uplink"
)

func PossibleValuesForDirection() []string {
	return []string{
		string(DirectionDownlink),
		string(DirectionUplink),
	}
}

type Polarization string

const (
	PolarizationLHCP             Polarization = "LHCP"
	PolarizationLinearHorizontal Polarization = "linearHorizontal"
	PolarizationLinearVertical   Polarization = "linearVertical"
	PolarizationRHCP             Polarization = "RHCP"
)

func PossibleValuesForPolarization() []string {
	return []string{
		string(PolarizationLHCP),
		string(PolarizationLinearHorizontal),
		string(PolarizationLinearVertical),
		string(PolarizationRHCP),
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
