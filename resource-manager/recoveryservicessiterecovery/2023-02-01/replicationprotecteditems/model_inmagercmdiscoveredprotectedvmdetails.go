package replicationprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageRcmDiscoveredProtectedVMDetails struct {
	CreatedTimestamp       *string   `json:"createdTimestamp,omitempty"`
	DataStores             *[]string `json:"datastores,omitempty"`
	IPAddresses            *[]string `json:"ipAddresses,omitempty"`
	IsDeleted              *bool     `json:"isDeleted,omitempty"`
	LastDiscoveryTimeInUtc *string   `json:"lastDiscoveryTimeInUtc,omitempty"`
	OsName                 *string   `json:"osName,omitempty"`
	PowerStatus            *string   `json:"powerStatus,omitempty"`
	UpdatedTimestamp       *string   `json:"updatedTimestamp,omitempty"`
	VCenterFqdn            *string   `json:"vCenterFqdn,omitempty"`
	VCenterId              *string   `json:"vCenterId,omitempty"`
	VMFqdn                 *string   `json:"vmFqdn,omitempty"`
	VMwareToolsStatus      *string   `json:"vmwareToolsStatus,omitempty"`
}
