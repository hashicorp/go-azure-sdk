package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeOperationPredicate struct {
	AffinityId            *string
	AllocationTime        *string
	IPAddress             *string
	Id                    *string
	IsDedicated           *bool
	LastBootTime          *string
	RunningTaskSlotsCount *int64
	RunningTasksCount     *int64
	StateTransitionTime   *string
	TotalTasksRun         *int64
	TotalTasksSucceeded   *int64
	Url                   *string
	VMSize                *string
}

func (p ComputeNodeOperationPredicate) Matches(input ComputeNode) bool {

	if p.AffinityId != nil && (input.AffinityId == nil || *p.AffinityId != *input.AffinityId) {
		return false
	}

	if p.AllocationTime != nil && (input.AllocationTime == nil || *p.AllocationTime != *input.AllocationTime) {
		return false
	}

	if p.IPAddress != nil && (input.IPAddress == nil || *p.IPAddress != *input.IPAddress) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDedicated != nil && (input.IsDedicated == nil || *p.IsDedicated != *input.IsDedicated) {
		return false
	}

	if p.LastBootTime != nil && (input.LastBootTime == nil || *p.LastBootTime != *input.LastBootTime) {
		return false
	}

	if p.RunningTaskSlotsCount != nil && (input.RunningTaskSlotsCount == nil || *p.RunningTaskSlotsCount != *input.RunningTaskSlotsCount) {
		return false
	}

	if p.RunningTasksCount != nil && (input.RunningTasksCount == nil || *p.RunningTasksCount != *input.RunningTasksCount) {
		return false
	}

	if p.StateTransitionTime != nil && (input.StateTransitionTime == nil || *p.StateTransitionTime != *input.StateTransitionTime) {
		return false
	}

	if p.TotalTasksRun != nil && (input.TotalTasksRun == nil || *p.TotalTasksRun != *input.TotalTasksRun) {
		return false
	}

	if p.TotalTasksSucceeded != nil && (input.TotalTasksSucceeded == nil || *p.TotalTasksSucceeded != *input.TotalTasksSucceeded) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	if p.VMSize != nil && (input.VMSize == nil || *p.VMSize != *input.VMSize) {
		return false
	}

	return true
}

type NodeVMExtensionOperationPredicate struct {
	ProvisioningState *string
}

func (p NodeVMExtensionOperationPredicate) Matches(input NodeVMExtension) bool {

	if p.ProvisioningState != nil && (input.ProvisioningState == nil || *p.ProvisioningState != *input.ProvisioningState) {
		return false
	}

	return true
}
