package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVMProfile struct {
	ApplicationProfile     *ApplicationProfile                     `json:"applicationProfile"`
	BillingProfile         *BillingProfile                         `json:"billingProfile"`
	CapacityReservation    *CapacityReservationProfile             `json:"capacityReservation"`
	DiagnosticsProfile     *DiagnosticsProfile                     `json:"diagnosticsProfile"`
	EvictionPolicy         *VirtualMachineEvictionPolicyTypes      `json:"evictionPolicy,omitempty"`
	ExtensionProfile       *VirtualMachineScaleSetExtensionProfile `json:"extensionProfile"`
	LicenseType            *string                                 `json:"licenseType,omitempty"`
	NetworkProfile         *VirtualMachineScaleSetNetworkProfile   `json:"networkProfile"`
	OsProfile              *VirtualMachineScaleSetOSProfile        `json:"osProfile"`
	Priority               *VirtualMachinePriorityTypes            `json:"priority,omitempty"`
	ScheduledEventsProfile *ScheduledEventsProfile                 `json:"scheduledEventsProfile"`
	SecurityProfile        *SecurityProfile                        `json:"securityProfile"`
	StorageProfile         *VirtualMachineScaleSetStorageProfile   `json:"storageProfile"`
	UserData               *string                                 `json:"userData,omitempty"`
}
