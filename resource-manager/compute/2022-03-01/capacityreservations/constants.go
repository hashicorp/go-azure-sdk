package capacityreservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationInstanceViewTypes string

const (
	CapacityReservationInstanceViewTypesInstanceView CapacityReservationInstanceViewTypes = "instanceView"
)

func PossibleValuesForCapacityReservationInstanceViewTypes() []string {
	return []string{
		string(CapacityReservationInstanceViewTypesInstanceView),
	}
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}
