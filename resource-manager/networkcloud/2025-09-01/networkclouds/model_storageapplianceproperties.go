package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageApplianceProperties struct {
	AdministratorCredentials      AdministrativeCredentials          `json:"administratorCredentials"`
	CaCertificate                 *CertificateInfo                   `json:"caCertificate,omitempty"`
	Capacity                      *int64                             `json:"capacity,omitempty"`
	CapacityUsed                  *int64                             `json:"capacityUsed,omitempty"`
	ClusterId                     *string                            `json:"clusterId,omitempty"`
	DetailedStatus                *StorageApplianceDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage         *string                            `json:"detailedStatusMessage,omitempty"`
	ManagementIPv4Address         *string                            `json:"managementIpv4Address,omitempty"`
	Manufacturer                  *string                            `json:"manufacturer,omitempty"`
	Model                         *string                            `json:"model,omitempty"`
	ProvisioningState             *StorageApplianceProvisioningState `json:"provisioningState,omitempty"`
	RackId                        string                             `json:"rackId"`
	RackSlot                      int64                              `json:"rackSlot"`
	RemoteVendorManagementFeature *RemoteVendorManagementFeature     `json:"remoteVendorManagementFeature,omitempty"`
	RemoteVendorManagementStatus  *RemoteVendorManagementStatus      `json:"remoteVendorManagementStatus,omitempty"`
	SecretRotationStatus          *[]SecretRotationStatus            `json:"secretRotationStatus,omitempty"`
	SerialNumber                  string                             `json:"serialNumber"`
	StorageApplianceSkuId         string                             `json:"storageApplianceSkuId"`
	Version                       *string                            `json:"version,omitempty"`
}
