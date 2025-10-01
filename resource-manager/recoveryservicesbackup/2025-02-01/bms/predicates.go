package bms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupManagementUsageOperationPredicate struct {
	CurrentValue  *int64
	Limit         *int64
	NextResetTime *string
	QuotaPeriod   *string
}

func (p BackupManagementUsageOperationPredicate) Matches(input BackupManagementUsage) bool {

	if p.CurrentValue != nil && (input.CurrentValue == nil || *p.CurrentValue != *input.CurrentValue) {
		return false
	}

	if p.Limit != nil && (input.Limit == nil || *p.Limit != *input.Limit) {
		return false
	}

	if p.NextResetTime != nil && (input.NextResetTime == nil || *p.NextResetTime != *input.NextResetTime) {
		return false
	}

	if p.QuotaPeriod != nil && (input.QuotaPeriod == nil || *p.QuotaPeriod != *input.QuotaPeriod) {
		return false
	}

	return true
}

type ProtectableContainerResourceOperationPredicate struct {
	ETag     *string
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ProtectableContainerResourceOperationPredicate) Matches(input ProtectableContainerResource) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ProtectedItemResourceOperationPredicate struct {
	ETag     *string
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ProtectedItemResourceOperationPredicate) Matches(input ProtectedItemResource) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ProtectionContainerResourceOperationPredicate struct {
	ETag     *string
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ProtectionContainerResourceOperationPredicate) Matches(input ProtectionContainerResource) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ProtectionIntentResourceOperationPredicate struct {
	ETag     *string
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ProtectionIntentResourceOperationPredicate) Matches(input ProtectionIntentResource) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type WorkloadProtectableItemResourceOperationPredicate struct {
	ETag     *string
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p WorkloadProtectableItemResourceOperationPredicate) Matches(input WorkloadProtectableItemResource) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
