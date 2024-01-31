package distributedavailabilitygroups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *DistributedAvailabilityGroupDatabase) GetLastBackupTimeAsTime() (*time.Time, error) {
	if o.LastBackupTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastBackupTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DistributedAvailabilityGroupDatabase) SetLastBackupTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastBackupTime = &formatted
}

func (o *DistributedAvailabilityGroupDatabase) GetLastCommitTimeAsTime() (*time.Time, error) {
	if o.LastCommitTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastCommitTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DistributedAvailabilityGroupDatabase) SetLastCommitTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastCommitTime = &formatted
}

func (o *DistributedAvailabilityGroupDatabase) GetLastHardenedTimeAsTime() (*time.Time, error) {
	if o.LastHardenedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastHardenedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DistributedAvailabilityGroupDatabase) SetLastHardenedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastHardenedTime = &formatted
}

func (o *DistributedAvailabilityGroupDatabase) GetLastReceivedTimeAsTime() (*time.Time, error) {
	if o.LastReceivedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastReceivedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DistributedAvailabilityGroupDatabase) SetLastReceivedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastReceivedTime = &formatted
}

func (o *DistributedAvailabilityGroupDatabase) GetLastSentTimeAsTime() (*time.Time, error) {
	if o.LastSentTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastSentTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DistributedAvailabilityGroupDatabase) SetLastSentTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastSentTime = &formatted
}
