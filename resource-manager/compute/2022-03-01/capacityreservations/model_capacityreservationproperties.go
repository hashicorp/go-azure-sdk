package capacityreservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationProperties struct {
	InstanceView              *CapacityReservationInstanceView `json:"instanceView,omitempty"`
	ProvisioningState         *string                          `json:"provisioningState,omitempty"`
	ProvisioningTime          *string                          `json:"provisioningTime,omitempty"`
	ReservationId             *string                          `json:"reservationId,omitempty"`
	TimeCreated               *string                          `json:"timeCreated,omitempty"`
	VirtualMachinesAssociated *[]SubResourceReadOnly           `json:"virtualMachinesAssociated,omitempty"`
}
