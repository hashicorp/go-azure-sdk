package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerSourceDataSetOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ConsumerSourceDataSetOperationPredicate) Matches(input ConsumerSourceDataSet) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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

type ShareSubscriptionOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ShareSubscriptionOperationPredicate) Matches(input ShareSubscription) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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

type ShareSubscriptionSynchronizationListOperationPredicate struct {
	NextLink *string
}

func (p ShareSubscriptionSynchronizationListOperationPredicate) Matches(input ShareSubscriptionSynchronizationList) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}

type SourceShareSynchronizationSettingListOperationPredicate struct {
	NextLink *string
}

func (p SourceShareSynchronizationSettingListOperationPredicate) Matches(input SourceShareSynchronizationSettingList) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}

type SynchronizationDetailsListOperationPredicate struct {
	NextLink *string
}

func (p SynchronizationDetailsListOperationPredicate) Matches(input SynchronizationDetailsList) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}
