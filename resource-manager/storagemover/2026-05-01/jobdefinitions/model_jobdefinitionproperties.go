package jobdefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobDefinitionProperties struct {
	AgentName                     *string                                 `json:"agentName,omitempty"`
	AgentResourceId               *string                                 `json:"agentResourceId,omitempty"`
	Connections                   *[]string                               `json:"connections,omitempty"`
	CopyMode                      CopyMode                                `json:"copyMode"`
	CrossTenantEndpointResourceId *string                                 `json:"crossTenantEndpointResourceId,omitempty"`
	CrossTenantEndpointTenantId   *string                                 `json:"crossTenantEndpointTenantId,omitempty"`
	DataIntegrityValidation       *DataIntegrityValidation                `json:"dataIntegrityValidation,omitempty"`
	Description                   *string                                 `json:"description,omitempty"`
	IsCrossTenantJob              *bool                                   `json:"isCrossTenantJob,omitempty"`
	JobType                       *JobType                                `json:"jobType,omitempty"`
	LatestJobRunName              *string                                 `json:"latestJobRunName,omitempty"`
	LatestJobRunResourceId        *string                                 `json:"latestJobRunResourceId,omitempty"`
	LatestJobRunStatus            *JobRunStatus                           `json:"latestJobRunStatus,omitempty"`
	MoverSyncedUntil              *string                                 `json:"moverSyncedUntil,omitempty"`
	PreservePermissions           *bool                                   `json:"preservePermissions,omitempty"`
	ProvisioningState             *ProvisioningState                      `json:"provisioningState,omitempty"`
	Schedule                      *ScheduleInfo                           `json:"schedule,omitempty"`
	SourceName                    string                                  `json:"sourceName"`
	SourceResourceId              *string                                 `json:"sourceResourceId,omitempty"`
	SourceSubpath                 *string                                 `json:"sourceSubpath,omitempty"`
	SourceTargetMap               *JobDefinitionPropertiesSourceTargetMap `json:"sourceTargetMap,omitempty"`
	SyncMode                      *string                                 `json:"syncMode,omitempty"`
	TargetName                    string                                  `json:"targetName"`
	TargetResourceId              *string                                 `json:"targetResourceId,omitempty"`
	TargetSubpath                 *string                                 `json:"targetSubpath,omitempty"`
}

func (o *JobDefinitionProperties) GetMoverSyncedUntilAsTime() (*time.Time, error) {
	if o.MoverSyncedUntil == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MoverSyncedUntil, "2006-01-02T15:04:05Z07:00")
}

func (o *JobDefinitionProperties) SetMoverSyncedUntilAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MoverSyncedUntil = &formatted
}
