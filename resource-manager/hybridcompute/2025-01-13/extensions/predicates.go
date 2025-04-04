package extensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionPublisherOperationPredicate struct {
	Id   *string
	Name *string
}

func (p ExtensionPublisherOperationPredicate) Matches(input ExtensionPublisher) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}

type ExtensionTypeOperationPredicate struct {
	Id   *string
	Name *string
}

func (p ExtensionTypeOperationPredicate) Matches(input ExtensionType) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}

type ExtensionValueV2OperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ExtensionValueV2OperationPredicate) Matches(input ExtensionValueV2) bool {

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
