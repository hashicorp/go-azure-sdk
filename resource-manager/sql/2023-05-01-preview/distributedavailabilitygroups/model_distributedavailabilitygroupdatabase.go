package distributedavailabilitygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributedAvailabilityGroupDatabase struct {
	ConnectedState                    *ReplicaConnectedState        `json:"connectedState,omitempty"`
	DatabaseName                      *string                       `json:"databaseName,omitempty"`
	InstanceRedoReplicationLagSeconds *int64                        `json:"instanceRedoReplicationLagSeconds,omitempty"`
	InstanceReplicaId                 *string                       `json:"instanceReplicaId,omitempty"`
	InstanceSendReplicationLagSeconds *int64                        `json:"instanceSendReplicationLagSeconds,omitempty"`
	LastBackupLsn                     *string                       `json:"lastBackupLsn,omitempty"`
	LastBackupTime                    *string                       `json:"lastBackupTime,omitempty"`
	LastCommitLsn                     *string                       `json:"lastCommitLsn,omitempty"`
	LastCommitTime                    *string                       `json:"lastCommitTime,omitempty"`
	LastHardenedLsn                   *string                       `json:"lastHardenedLsn,omitempty"`
	LastHardenedTime                  *string                       `json:"lastHardenedTime,omitempty"`
	LastReceivedLsn                   *string                       `json:"lastReceivedLsn,omitempty"`
	LastReceivedTime                  *string                       `json:"lastReceivedTime,omitempty"`
	LastSentLsn                       *string                       `json:"lastSentLsn,omitempty"`
	LastSentTime                      *string                       `json:"lastSentTime,omitempty"`
	MostRecentLinkError               *string                       `json:"mostRecentLinkError,omitempty"`
	PartnerAuthCertValidity           *CertificateInfo              `json:"partnerAuthCertValidity,omitempty"`
	PartnerReplicaId                  *string                       `json:"partnerReplicaId,omitempty"`
	ReplicaState                      *string                       `json:"replicaState,omitempty"`
	SynchronizationHealth             *ReplicaSynchronizationHealth `json:"synchronizationHealth,omitempty"`
}
