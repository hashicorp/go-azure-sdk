package groundstation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityParameter string

const (
	CapabilityParameterCommunication    CapabilityParameter = "Communication"
	CapabilityParameterEarthObservation CapabilityParameter = "EarthObservation"
)

func PossibleValuesForCapabilityParameter() []string {
	return []string{
		string(CapabilityParameterCommunication),
		string(CapabilityParameterEarthObservation),
	}
}

type ReleaseMode string

const (
	ReleaseModeGA      ReleaseMode = "GA"
	ReleaseModePreview ReleaseMode = "Preview"
)

func PossibleValuesForReleaseMode() []string {
	return []string{
		string(ReleaseModeGA),
		string(ReleaseModePreview),
	}
}
