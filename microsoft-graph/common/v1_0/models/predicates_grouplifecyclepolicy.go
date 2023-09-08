package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupLifecyclePolicyOperationPredicate struct {
	AlternateNotificationEmails *string
	GroupLifetimeInDays         *int64
	Id                          *string
	ManagedGroupTypes           *string
	ODataType                   *string
}

func (p GroupLifecyclePolicyOperationPredicate) Matches(input GroupLifecyclePolicy) bool {

	if p.AlternateNotificationEmails != nil && (input.AlternateNotificationEmails == nil || *p.AlternateNotificationEmails != *input.AlternateNotificationEmails) {
		return false
	}

	if p.GroupLifetimeInDays != nil && (input.GroupLifetimeInDays == nil || *p.GroupLifetimeInDays != *input.GroupLifetimeInDays) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedGroupTypes != nil && (input.ManagedGroupTypes == nil || *p.ManagedGroupTypes != *input.ManagedGroupTypes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
