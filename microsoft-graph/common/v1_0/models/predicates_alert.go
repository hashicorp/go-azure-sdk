package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertOperationPredicate struct {
	ActivityGroupName    *string
	AssignedTo           *string
	AzureSubscriptionId  *string
	AzureTenantId        *string
	Category             *string
	ClosedDateTime       *string
	Confidence           *int64
	CreatedDateTime      *string
	Description          *string
	EventDateTime        *string
	Id                   *string
	LastEventDateTime    *string
	LastModifiedDateTime *string
	ODataType            *string
	Title                *string
}

func (p AlertOperationPredicate) Matches(input Alert) bool {

	if p.ActivityGroupName != nil && (input.ActivityGroupName == nil || *p.ActivityGroupName != *input.ActivityGroupName) {
		return false
	}

	if p.AssignedTo != nil && (input.AssignedTo == nil || *p.AssignedTo != *input.AssignedTo) {
		return false
	}

	if p.AzureSubscriptionId != nil && (input.AzureSubscriptionId == nil || *p.AzureSubscriptionId != *input.AzureSubscriptionId) {
		return false
	}

	if p.AzureTenantId != nil && (input.AzureTenantId == nil || *p.AzureTenantId != *input.AzureTenantId) {
		return false
	}

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.ClosedDateTime != nil && (input.ClosedDateTime == nil || *p.ClosedDateTime != *input.ClosedDateTime) {
		return false
	}

	if p.Confidence != nil && (input.Confidence == nil || *p.Confidence != *input.Confidence) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastEventDateTime != nil && (input.LastEventDateTime == nil || *p.LastEventDateTime != *input.LastEventDateTime) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
