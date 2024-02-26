package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SasTokenInformationListResultOperationPredicate struct {
	NextLink *string
}

func (p SasTokenInformationListResultOperationPredicate) Matches(input SasTokenInformationListResult) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}

type StorageAccountInformationOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p StorageAccountInformationOperationPredicate) Matches(input StorageAccountInformation) bool {

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

type StorageContainerOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p StorageContainerOperationPredicate) Matches(input StorageContainer) bool {

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
