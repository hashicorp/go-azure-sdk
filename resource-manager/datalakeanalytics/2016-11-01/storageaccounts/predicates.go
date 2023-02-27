package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SasTokenInformationOperationPredicate struct {
	AccessToken *string
}

func (p SasTokenInformationOperationPredicate) Matches(input SasTokenInformation) bool {

	if p.AccessToken != nil && (input.AccessToken == nil && *p.AccessToken != *input.AccessToken) {
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

type StorageContainerOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p StorageContainerOperationPredicate) Matches(input StorageContainer) bool {

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
