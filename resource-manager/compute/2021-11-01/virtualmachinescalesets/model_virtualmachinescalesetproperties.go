package virtualmachinescalesets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetProperties struct {
	AdditionalCapabilities                 *AdditionalCapabilities          `json:"additionalCapabilities"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy          `json:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                            `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
	HostGroup                              *SubResource                     `json:"hostGroup"`
	OrchestrationMode                      *OrchestrationMode               `json:"orchestrationMode,omitempty"`
	Overprovision                          *bool                            `json:"overprovision,omitempty"`
	PlatformFaultDomainCount               *int64                           `json:"platformFaultDomainCount,omitempty"`
	ProvisioningState                      *string                          `json:"provisioningState,omitempty"`
	ProximityPlacementGroup                *SubResource                     `json:"proximityPlacementGroup"`
	ScaleInPolicy                          *ScaleInPolicy                   `json:"scaleInPolicy"`
	SinglePlacementGroup                   *bool                            `json:"singlePlacementGroup,omitempty"`
	SpotRestorePolicy                      *SpotRestorePolicy               `json:"spotRestorePolicy"`
	TimeCreated                            *string                          `json:"timeCreated,omitempty"`
	UniqueId                               *string                          `json:"uniqueId,omitempty"`
	UpgradePolicy                          *UpgradePolicy                   `json:"upgradePolicy"`
	VirtualMachineProfile                  *VirtualMachineScaleSetVMProfile `json:"virtualMachineProfile"`
	ZoneBalance                            *bool                            `json:"zoneBalance,omitempty"`
}

func (o *VirtualMachineScaleSetProperties) GetTimeCreatedAsTime() (*time.Time, error) {
	if o.TimeCreated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeCreated, "2006-01-02T15:04:05Z07:00")
}

func (o *VirtualMachineScaleSetProperties) SetTimeCreatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeCreated = &formatted
}
