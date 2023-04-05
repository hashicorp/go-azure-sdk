package capacityreservationgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationGroupInstanceViewTypes string

const (
	CapacityReservationGroupInstanceViewTypesInstanceView CapacityReservationGroupInstanceViewTypes = "instanceView"
)

func PossibleValuesForCapacityReservationGroupInstanceViewTypes() []string {
	return []string{
		string(CapacityReservationGroupInstanceViewTypesInstanceView),
	}
}

type ExpandTypesForGetCapacityReservationGroups string

const (
	ExpandTypesForGetCapacityReservationGroupsVirtualMachineScaleSetVMsRef ExpandTypesForGetCapacityReservationGroups = "virtualMachineScaleSetVMs/$ref"
	ExpandTypesForGetCapacityReservationGroupsVirtualMachinesRef           ExpandTypesForGetCapacityReservationGroups = "virtualMachines/$ref"
)

func PossibleValuesForExpandTypesForGetCapacityReservationGroups() []string {
	return []string{
		string(ExpandTypesForGetCapacityReservationGroupsVirtualMachineScaleSetVMsRef),
		string(ExpandTypesForGetCapacityReservationGroupsVirtualMachinesRef),
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
