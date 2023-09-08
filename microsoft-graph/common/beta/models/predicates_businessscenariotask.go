package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskOperationPredicate struct {
	ActiveChecklistItemCount *int64
	AssigneePriority         *string
	BucketId                 *string
	ChecklistItemCount       *int64
	CompletedDateTime        *string
	ConversationThreadId     *string
	CreatedDateTime          *string
	DueDateTime              *string
	HasDescription           *bool
	Id                       *string
	ODataType                *string
	OrderHint                *string
	PercentComplete          *int64
	PlanId                   *string
	Priority                 *int64
	ReferenceCount           *int64
	StartDateTime            *string
	Title                    *string
}

func (p BusinessScenarioTaskOperationPredicate) Matches(input BusinessScenarioTask) bool {

	if p.ActiveChecklistItemCount != nil && (input.ActiveChecklistItemCount == nil || *p.ActiveChecklistItemCount != *input.ActiveChecklistItemCount) {
		return false
	}

	if p.AssigneePriority != nil && (input.AssigneePriority == nil || *p.AssigneePriority != *input.AssigneePriority) {
		return false
	}

	if p.BucketId != nil && (input.BucketId == nil || *p.BucketId != *input.BucketId) {
		return false
	}

	if p.ChecklistItemCount != nil && (input.ChecklistItemCount == nil || *p.ChecklistItemCount != *input.ChecklistItemCount) {
		return false
	}

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.ConversationThreadId != nil && (input.ConversationThreadId == nil || *p.ConversationThreadId != *input.ConversationThreadId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DueDateTime != nil && (input.DueDateTime == nil || *p.DueDateTime != *input.DueDateTime) {
		return false
	}

	if p.HasDescription != nil && (input.HasDescription == nil || *p.HasDescription != *input.HasDescription) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrderHint != nil && (input.OrderHint == nil || *p.OrderHint != *input.OrderHint) {
		return false
	}

	if p.PercentComplete != nil && (input.PercentComplete == nil || *p.PercentComplete != *input.PercentComplete) {
		return false
	}

	if p.PlanId != nil && (input.PlanId == nil || *p.PlanId != *input.PlanId) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.ReferenceCount != nil && (input.ReferenceCount == nil || *p.ReferenceCount != *input.ReferenceCount) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
