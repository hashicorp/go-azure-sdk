package syncmembers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncFullSchemaPropertiesOperationPredicate struct {
	LastUpdateTime *string
}

func (p SyncFullSchemaPropertiesOperationPredicate) Matches(input SyncFullSchemaProperties) bool {

	if p.LastUpdateTime != nil && (input.LastUpdateTime == nil && *p.LastUpdateTime != *input.LastUpdateTime) {
		return false
	}

	return true
}

type SyncMemberOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p SyncMemberOperationPredicate) Matches(input SyncMember) bool {

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
