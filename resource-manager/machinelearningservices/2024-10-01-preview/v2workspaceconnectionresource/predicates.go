package v2workspaceconnectionresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate) Matches(input EndpointDeploymentResourcePropertiesBasicResource) bool {

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

type EndpointModelPropertiesOperationPredicate struct {
	Format           *string
	IsDefaultVersion *bool
	MaxCapacity      *int64
	Name             *string
	Version          *string
}

func (p EndpointModelPropertiesOperationPredicate) Matches(input EndpointModelProperties) bool {

	if p.Format != nil && (input.Format == nil || *p.Format != *input.Format) {
		return false
	}

	if p.IsDefaultVersion != nil && (input.IsDefaultVersion == nil || *p.IsDefaultVersion != *input.IsDefaultVersion) {
		return false
	}

	if p.MaxCapacity != nil && (input.MaxCapacity == nil || *p.MaxCapacity != *input.MaxCapacity) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}

type RaiBlocklistItemPropertiesBasicResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p RaiBlocklistItemPropertiesBasicResourceOperationPredicate) Matches(input RaiBlocklistItemPropertiesBasicResource) bool {

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

type RaiBlocklistPropertiesBasicResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p RaiBlocklistPropertiesBasicResourceOperationPredicate) Matches(input RaiBlocklistPropertiesBasicResource) bool {

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

type RaiPolicyPropertiesBasicResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p RaiPolicyPropertiesBasicResourceOperationPredicate) Matches(input RaiPolicyPropertiesBasicResource) bool {

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

type WorkspaceConnectionPropertiesV2BasicResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkspaceConnectionPropertiesV2BasicResourceOperationPredicate) Matches(input WorkspaceConnectionPropertiesV2BasicResource) bool {

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
