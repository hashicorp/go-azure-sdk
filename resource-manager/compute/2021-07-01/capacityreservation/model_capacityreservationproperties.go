package capacityreservation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationProperties struct {
	InstanceView              *CapacityReservationInstanceView `json:"instanceView,omitempty"`
	ProvisioningState         *string                          `json:"provisioningState,omitempty"`
	ProvisioningTime          *string                          `json:"provisioningTime,omitempty"`
	ReservationId             *string                          `json:"reservationId,omitempty"`
	VirtualMachinesAssociated *[]SubResourceReadOnly           `json:"virtualMachinesAssociated,omitempty"`
}

func (o *CapacityReservationProperties) GetProvisioningTimeAsTime() (*time.Time, error) {
	if o.ProvisioningTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ProvisioningTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CapacityReservationProperties) SetProvisioningTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ProvisioningTime = &formatted
}
