package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsOrgOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p AzureDevOpsOrgOperationPredicate) Matches(input AzureDevOpsOrg) bool {

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

type AzureDevOpsProjectOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p AzureDevOpsProjectOperationPredicate) Matches(input AzureDevOpsProject) bool {

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

type AzureDevOpsRepositoryOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p AzureDevOpsRepositoryOperationPredicate) Matches(input AzureDevOpsRepository) bool {

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

type DevOpsConfigurationOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p DevOpsConfigurationOperationPredicate) Matches(input DevOpsConfiguration) bool {

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

type GitHubOwnerOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GitHubOwnerOperationPredicate) Matches(input GitHubOwner) bool {

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

type GitHubRepositoryOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GitHubRepositoryOperationPredicate) Matches(input GitHubRepository) bool {

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

type GitLabGroupOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GitLabGroupOperationPredicate) Matches(input GitLabGroup) bool {

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

type GitLabProjectOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GitLabProjectOperationPredicate) Matches(input GitLabProject) bool {

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
