package pools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPool struct {
	AllocationState               *AllocationState               `json:"allocationState,omitempty"`
	AllocationStateTransitionTime *string                        `json:"allocationStateTransitionTime,omitempty"`
	ApplicationLicenses           *[]string                      `json:"applicationLicenses,omitempty"`
	ApplicationPackageReferences  *[]ApplicationPackageReference `json:"applicationPackageReferences,omitempty"`
	AutoScaleEvaluationInterval   *string                        `json:"autoScaleEvaluationInterval,omitempty"`
	AutoScaleFormula              *string                        `json:"autoScaleFormula,omitempty"`
	AutoScaleRun                  *AutoScaleRun                  `json:"autoScaleRun,omitempty"`
	CertificateReferences         *[]CertificateReference        `json:"certificateReferences,omitempty"`
	CloudServiceConfiguration     *CloudServiceConfiguration     `json:"cloudServiceConfiguration,omitempty"`
	CreationTime                  *string                        `json:"creationTime,omitempty"`
	CurrentDedicatedNodes         *int64                         `json:"currentDedicatedNodes,omitempty"`
	CurrentLowPriorityNodes       *int64                         `json:"currentLowPriorityNodes,omitempty"`
	DisplayName                   *string                        `json:"displayName,omitempty"`
	ETag                          *string                        `json:"eTag,omitempty"`
	EnableAutoScale               *bool                          `json:"enableAutoScale,omitempty"`
	EnableInterNodeCommunication  *bool                          `json:"enableInterNodeCommunication,omitempty"`
	Id                            *string                        `json:"id,omitempty"`
	Identity                      *BatchPoolIdentity             `json:"identity,omitempty"`
	LastModified                  *string                        `json:"lastModified,omitempty"`
	Metadata                      *[]MetadataItem                `json:"metadata,omitempty"`
	MountConfiguration            *[]MountConfiguration          `json:"mountConfiguration,omitempty"`
	NetworkConfiguration          *NetworkConfiguration          `json:"networkConfiguration,omitempty"`
	ResizeErrors                  *[]ResizeError                 `json:"resizeErrors,omitempty"`
	ResizeTimeout                 *string                        `json:"resizeTimeout,omitempty"`
	StartTask                     *StartTask                     `json:"startTask,omitempty"`
	State                         *PoolState                     `json:"state,omitempty"`
	StateTransitionTime           *string                        `json:"stateTransitionTime,omitempty"`
	Stats                         *PoolStatistics                `json:"stats,omitempty"`
	TargetDedicatedNodes          *int64                         `json:"targetDedicatedNodes,omitempty"`
	TargetLowPriorityNodes        *int64                         `json:"targetLowPriorityNodes,omitempty"`
	TaskSchedulingPolicy          *TaskSchedulingPolicy          `json:"taskSchedulingPolicy,omitempty"`
	TaskSlotsPerNode              *int64                         `json:"taskSlotsPerNode,omitempty"`
	Url                           *string                        `json:"url,omitempty"`
	UserAccounts                  *[]UserAccount                 `json:"userAccounts,omitempty"`
	VMSize                        *string                        `json:"vmSize,omitempty"`
	VirtualMachineConfiguration   *VirtualMachineConfiguration   `json:"virtualMachineConfiguration,omitempty"`
}

func (o *CloudPool) GetAllocationStateTransitionTimeAsTime() (*time.Time, error) {
	if o.AllocationStateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AllocationStateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudPool) SetAllocationStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AllocationStateTransitionTime = &formatted
}

func (o *CloudPool) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudPool) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}

func (o *CloudPool) GetLastModifiedAsTime() (*time.Time, error) {
	if o.LastModified == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModified, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudPool) SetLastModifiedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModified = &formatted
}

func (o *CloudPool) GetStateTransitionTimeAsTime() (*time.Time, error) {
	if o.StateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudPool) SetStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StateTransitionTime = &formatted
}
