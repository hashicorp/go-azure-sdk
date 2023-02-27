package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerSourceDataSetOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ConsumerSourceDataSetOperationPredicate) Matches(input ConsumerSourceDataSet) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
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

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type ShareSubscriptionSynchronizationOperationPredicate struct {
	DurationMs        *int64
	EndTime           *string
	Message           *string
	StartTime         *string
	Status            *string
	SynchronizationId *string
}

func (p ShareSubscriptionSynchronizationOperationPredicate) Matches(input ShareSubscriptionSynchronization) bool {

	if p.DurationMs != nil && (input.DurationMs == nil && *p.DurationMs != *input.DurationMs) {
		return false
	}

	if p.EndTime != nil && (input.EndTime == nil && *p.EndTime != *input.EndTime) {
		return false
	}

	if p.Message != nil && (input.Message == nil && *p.Message != *input.Message) {
		return false
	}

	if p.StartTime != nil && (input.StartTime == nil && *p.StartTime != *input.StartTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil && *p.Status != *input.Status) {
		return false
	}

	if p.SynchronizationId != nil && *p.SynchronizationId != input.SynchronizationId {
		return false
	}

	return true
}

type SourceShareSynchronizationSettingOperationPredicate struct {
}

func (p SourceShareSynchronizationSettingOperationPredicate) Matches(input SourceShareSynchronizationSetting) bool {

	return true
}

type SynchronizationDetailsOperationPredicate struct {
	DataSetId    *string
	DurationMs   *int64
	EndTime      *string
	FilesRead    *int64
	FilesWritten *int64
	Message      *string
	Name         *string
	RowsCopied   *int64
	RowsRead     *int64
	SizeRead     *int64
	SizeWritten  *int64
	StartTime    *string
	Status       *string
	VCore        *int64
}

func (p SynchronizationDetailsOperationPredicate) Matches(input SynchronizationDetails) bool {

	if p.DataSetId != nil && (input.DataSetId == nil && *p.DataSetId != *input.DataSetId) {
		return false
	}

	if p.DurationMs != nil && (input.DurationMs == nil && *p.DurationMs != *input.DurationMs) {
		return false
	}

	if p.EndTime != nil && (input.EndTime == nil && *p.EndTime != *input.EndTime) {
		return false
	}

	if p.FilesRead != nil && (input.FilesRead == nil && *p.FilesRead != *input.FilesRead) {
		return false
	}

	if p.FilesWritten != nil && (input.FilesWritten == nil && *p.FilesWritten != *input.FilesWritten) {
		return false
	}

	if p.Message != nil && (input.Message == nil && *p.Message != *input.Message) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.RowsCopied != nil && (input.RowsCopied == nil && *p.RowsCopied != *input.RowsCopied) {
		return false
	}

	if p.RowsRead != nil && (input.RowsRead == nil && *p.RowsRead != *input.RowsRead) {
		return false
	}

	if p.SizeRead != nil && (input.SizeRead == nil && *p.SizeRead != *input.SizeRead) {
		return false
	}

	if p.SizeWritten != nil && (input.SizeWritten == nil && *p.SizeWritten != *input.SizeWritten) {
		return false
	}

	if p.StartTime != nil && (input.StartTime == nil && *p.StartTime != *input.StartTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil && *p.Status != *input.Status) {
		return false
	}

	if p.VCore != nil && (input.VCore == nil && *p.VCore != *input.VCore) {
		return false
	}

	return true
}
