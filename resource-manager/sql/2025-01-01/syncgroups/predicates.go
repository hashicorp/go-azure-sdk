package syncgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncDatabaseIdPropertiesOperationPredicate struct {
	Id *string
}

func (p SyncDatabaseIdPropertiesOperationPredicate) Matches(input SyncDatabaseIdProperties) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}

type SyncFullSchemaPropertiesOperationPredicate struct {
	LastUpdateTime *string
}

func (p SyncFullSchemaPropertiesOperationPredicate) Matches(input SyncFullSchemaProperties) bool {

	if p.LastUpdateTime != nil && (input.LastUpdateTime == nil || *p.LastUpdateTime != *input.LastUpdateTime) {
		return false
	}

	return true
}

type SyncGroupOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p SyncGroupOperationPredicate) Matches(input SyncGroup) bool {

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

type SyncGroupLogPropertiesOperationPredicate struct {
	Details         *string
	OperationStatus *string
	Source          *string
	Timestamp       *string
	TracingId       *string
}

func (p SyncGroupLogPropertiesOperationPredicate) Matches(input SyncGroupLogProperties) bool {

	if p.Details != nil && (input.Details == nil || *p.Details != *input.Details) {
		return false
	}

	if p.OperationStatus != nil && (input.OperationStatus == nil || *p.OperationStatus != *input.OperationStatus) {
		return false
	}

	if p.Source != nil && (input.Source == nil || *p.Source != *input.Source) {
		return false
	}

	if p.Timestamp != nil && (input.Timestamp == nil || *p.Timestamp != *input.Timestamp) {
		return false
	}

	if p.TracingId != nil && (input.TracingId == nil || *p.TracingId != *input.TracingId) {
		return false
	}

	return true
}
