package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVMProfile struct {
	ApplicationProfile     *ApplicationProfile                     `json:"applicationProfile,omitempty"`
	BillingProfile         *BillingProfile                         `json:"billingProfile,omitempty"`
	CapacityReservation    *CapacityReservationProfile             `json:"capacityReservation,omitempty"`
	DiagnosticsProfile     *DiagnosticsProfile                     `json:"diagnosticsProfile,omitempty"`
	EvictionPolicy         *VirtualMachineEvictionPolicyTypes      `json:"evictionPolicy,omitempty"`
	ExtensionProfile       *VirtualMachineScaleSetExtensionProfile `json:"extensionProfile,omitempty"`
	LicenseType            *string                                 `json:"licenseType,omitempty"`
	NetworkProfile         *VirtualMachineScaleSetNetworkProfile   `json:"networkProfile,omitempty"`
	OsProfile              *VirtualMachineScaleSetOSProfile        `json:"osProfile,omitempty"`
	Priority               *VirtualMachinePriorityTypes            `json:"priority,omitempty"`
	ScheduledEventsProfile *ScheduledEventsProfile                 `json:"scheduledEventsProfile,omitempty"`
	SecurityProfile        *SecurityProfile                        `json:"securityProfile,omitempty"`
	StorageProfile         *VirtualMachineScaleSetStorageProfile   `json:"storageProfile,omitempty"`
	UserData               *string                                 `json:"userData,omitempty"`
}
