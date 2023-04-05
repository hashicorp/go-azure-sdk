package contactprofile

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoTrackingConfiguration string

const (
	AutoTrackingConfigurationDisabled AutoTrackingConfiguration = "disabled"
	AutoTrackingConfigurationSBand    AutoTrackingConfiguration = "sBand"
	AutoTrackingConfigurationXBand    AutoTrackingConfiguration = "xBand"
)

func PossibleValuesForAutoTrackingConfiguration() []string {
	return []string{
		string(AutoTrackingConfigurationDisabled),
		string(AutoTrackingConfigurationSBand),
		string(AutoTrackingConfigurationXBand),
	}
}

type Direction string

const (
	DirectionDownlink Direction = "Downlink"
	DirectionUplink   Direction = "Uplink"
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

type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolTCP),
		string(ProtocolUDP),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "canceled"
	ProvisioningStateCreating  ProvisioningState = "creating"
	ProvisioningStateDeleting  ProvisioningState = "deleting"
	ProvisioningStateFailed    ProvisioningState = "failed"
	ProvisioningStateSucceeded ProvisioningState = "succeeded"
	ProvisioningStateUpdating  ProvisioningState = "updating"
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
