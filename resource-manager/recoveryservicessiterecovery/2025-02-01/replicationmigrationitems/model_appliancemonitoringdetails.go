package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplianceMonitoringDetails struct {
	CpuDetails              *ApplianceResourceDetails      `json:"cpuDetails,omitempty"`
	DatastoreSnapshot       *[]DataStoreUtilizationDetails `json:"datastoreSnapshot,omitempty"`
	DisksReplicationDetails *ApplianceResourceDetails      `json:"disksReplicationDetails,omitempty"`
	EsxiNfcBuffer           *ApplianceResourceDetails      `json:"esxiNfcBuffer,omitempty"`
	NetworkBandwidth        *ApplianceResourceDetails      `json:"networkBandwidth,omitempty"`
	RamDetails              *ApplianceResourceDetails      `json:"ramDetails,omitempty"`
}
