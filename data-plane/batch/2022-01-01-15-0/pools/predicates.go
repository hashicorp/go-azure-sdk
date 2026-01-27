package pools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPoolOperationPredicate struct {
	AllocationStateTransitionTime *string
	AutoScaleEvaluationInterval   *string
	AutoScaleFormula              *string
	CreationTime                  *string
	CurrentDedicatedNodes         *int64
	CurrentLowPriorityNodes       *int64
	DisplayName                   *string
	ETag                          *string
	EnableAutoScale               *bool
	EnableInterNodeCommunication  *bool
	Id                            *string
	LastModified                  *string
	ResizeTimeout                 *string
	StateTransitionTime           *string
	TargetDedicatedNodes          *int64
	TargetLowPriorityNodes        *int64
	TaskSlotsPerNode              *int64
	Url                           *string
	VMSize                        *string
}

func (p CloudPoolOperationPredicate) Matches(input CloudPool) bool {

	if p.AllocationStateTransitionTime != nil && (input.AllocationStateTransitionTime == nil || *p.AllocationStateTransitionTime != *input.AllocationStateTransitionTime) {
		return false
	}

	if p.AutoScaleEvaluationInterval != nil && (input.AutoScaleEvaluationInterval == nil || *p.AutoScaleEvaluationInterval != *input.AutoScaleEvaluationInterval) {
		return false
	}

	if p.AutoScaleFormula != nil && (input.AutoScaleFormula == nil || *p.AutoScaleFormula != *input.AutoScaleFormula) {
		return false
	}

	if p.CreationTime != nil && (input.CreationTime == nil || *p.CreationTime != *input.CreationTime) {
		return false
	}

	if p.CurrentDedicatedNodes != nil && (input.CurrentDedicatedNodes == nil || *p.CurrentDedicatedNodes != *input.CurrentDedicatedNodes) {
		return false
	}

	if p.CurrentLowPriorityNodes != nil && (input.CurrentLowPriorityNodes == nil || *p.CurrentLowPriorityNodes != *input.CurrentLowPriorityNodes) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.EnableAutoScale != nil && (input.EnableAutoScale == nil || *p.EnableAutoScale != *input.EnableAutoScale) {
		return false
	}

	if p.EnableInterNodeCommunication != nil && (input.EnableInterNodeCommunication == nil || *p.EnableInterNodeCommunication != *input.EnableInterNodeCommunication) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModified != nil && (input.LastModified == nil || *p.LastModified != *input.LastModified) {
		return false
	}

	if p.ResizeTimeout != nil && (input.ResizeTimeout == nil || *p.ResizeTimeout != *input.ResizeTimeout) {
		return false
	}

	if p.StateTransitionTime != nil && (input.StateTransitionTime == nil || *p.StateTransitionTime != *input.StateTransitionTime) {
		return false
	}

	if p.TargetDedicatedNodes != nil && (input.TargetDedicatedNodes == nil || *p.TargetDedicatedNodes != *input.TargetDedicatedNodes) {
		return false
	}

	if p.TargetLowPriorityNodes != nil && (input.TargetLowPriorityNodes == nil || *p.TargetLowPriorityNodes != *input.TargetLowPriorityNodes) {
		return false
	}

	if p.TaskSlotsPerNode != nil && (input.TaskSlotsPerNode == nil || *p.TaskSlotsPerNode != *input.TaskSlotsPerNode) {
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

type PoolUsageMetricsOperationPredicate struct {
	EndTime        *string
	PoolId         *string
	StartTime      *string
	TotalCoreHours *float64
	VMSize         *string
}

func (p PoolUsageMetricsOperationPredicate) Matches(input PoolUsageMetrics) bool {

	if p.EndTime != nil && *p.EndTime != input.EndTime {
		return false
	}

	if p.PoolId != nil && *p.PoolId != input.PoolId {
		return false
	}

	if p.StartTime != nil && *p.StartTime != input.StartTime {
		return false
	}

	if p.TotalCoreHours != nil && *p.TotalCoreHours != input.TotalCoreHours {
		return false
	}

	if p.VMSize != nil && *p.VMSize != input.VMSize {
		return false
	}

	return true
}
