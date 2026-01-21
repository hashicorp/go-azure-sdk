package computenodes

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNode struct {
	AffinityId            *string                           `json:"affinityId,omitempty"`
	AllocationTime        *string                           `json:"allocationTime,omitempty"`
	CertificateReferences *[]CertificateReference           `json:"certificateReferences,omitempty"`
	EndpointConfiguration *ComputeNodeEndpointConfiguration `json:"endpointConfiguration,omitempty"`
	Errors                *[]ComputeNodeError               `json:"errors,omitempty"`
	IPAddress             *string                           `json:"ipAddress,omitempty"`
	Id                    *string                           `json:"id,omitempty"`
	IsDedicated           *bool                             `json:"isDedicated,omitempty"`
	LastBootTime          *string                           `json:"lastBootTime,omitempty"`
	NodeAgentInfo         *NodeAgentInformation             `json:"nodeAgentInfo,omitempty"`
	RecentTasks           *[]TaskInformation                `json:"recentTasks,omitempty"`
	RunningTaskSlotsCount *int64                            `json:"runningTaskSlotsCount,omitempty"`
	RunningTasksCount     *int64                            `json:"runningTasksCount,omitempty"`
	SchedulingState       *SchedulingState                  `json:"schedulingState,omitempty"`
	StartTask             *StartTask                        `json:"startTask,omitempty"`
	StartTaskInfo         *StartTaskInformation             `json:"startTaskInfo,omitempty"`
	State                 *ComputeNodeState                 `json:"state,omitempty"`
	StateTransitionTime   *string                           `json:"stateTransitionTime,omitempty"`
	TotalTasksRun         *int64                            `json:"totalTasksRun,omitempty"`
	TotalTasksSucceeded   *int64                            `json:"totalTasksSucceeded,omitempty"`
	Url                   *string                           `json:"url,omitempty"`
	VMSize                *string                           `json:"vmSize,omitempty"`
	VirtualMachineInfo    *VirtualMachineInfo               `json:"virtualMachineInfo,omitempty"`
}

func (o *ComputeNode) GetAllocationTimeAsTime() (*time.Time, error) {
	if o.AllocationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AllocationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ComputeNode) SetAllocationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AllocationTime = &formatted
}

func (o *ComputeNode) GetLastBootTimeAsTime() (*time.Time, error) {
	if o.LastBootTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastBootTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ComputeNode) SetLastBootTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastBootTime = &formatted
}

func (o *ComputeNode) GetStateTransitionTimeAsTime() (*time.Time, error) {
	if o.StateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ComputeNode) SetStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StateTransitionTime = &formatted
}
