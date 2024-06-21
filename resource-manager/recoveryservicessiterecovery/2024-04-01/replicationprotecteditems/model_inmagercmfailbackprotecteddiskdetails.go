package replicationprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageRcmFailbackProtectedDiskDetails struct {
	CapacityInBytes               *int64                        `json:"capacityInBytes,omitempty"`
	DataPendingAtSourceAgentInMB  *float64                      `json:"dataPendingAtSourceAgentInMB,omitempty"`
	DataPendingInLogDataStoreInMB *float64                      `json:"dataPendingInLogDataStoreInMB,omitempty"`
	DiskId                        *string                       `json:"diskId,omitempty"`
	DiskName                      *string                       `json:"diskName,omitempty"`
	DiskUuid                      *string                       `json:"diskUuid,omitempty"`
	IrDetails                     *InMageRcmFailbackSyncDetails `json:"irDetails,omitempty"`
	IsInitialReplicationComplete  *string                       `json:"isInitialReplicationComplete,omitempty"`
	IsOSDisk                      *string                       `json:"isOSDisk,omitempty"`
	LastSyncTime                  *string                       `json:"lastSyncTime,omitempty"`
	ResyncDetails                 *InMageRcmFailbackSyncDetails `json:"resyncDetails,omitempty"`
}
