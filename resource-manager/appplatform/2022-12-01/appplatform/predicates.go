// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appplatform

type ApiPortalCustomDomainResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ApiPortalCustomDomainResourceOperationPredicate) Matches(input ApiPortalCustomDomainResource) bool {

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

type ApiPortalResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ApiPortalResourceOperationPredicate) Matches(input ApiPortalResource) bool {

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

type AppResourceOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p AppResourceOperationPredicate) Matches(input AppResource) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
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

type BindingResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BindingResourceOperationPredicate) Matches(input BindingResource) bool {

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

type BuildOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BuildOperationPredicate) Matches(input Build) bool {

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

type BuildResultOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BuildResultOperationPredicate) Matches(input BuildResult) bool {

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

type BuildServiceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BuildServiceOperationPredicate) Matches(input BuildService) bool {

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

type BuildServiceAgentPoolResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BuildServiceAgentPoolResourceOperationPredicate) Matches(input BuildServiceAgentPoolResource) bool {

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

type BuilderResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BuilderResourceOperationPredicate) Matches(input BuilderResource) bool {

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

type BuildpackBindingResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BuildpackBindingResourceOperationPredicate) Matches(input BuildpackBindingResource) bool {

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

type CertificateResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p CertificateResourceOperationPredicate) Matches(input CertificateResource) bool {

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

type ConfigurationServiceResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ConfigurationServiceResourceOperationPredicate) Matches(input ConfigurationServiceResource) bool {

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

type CustomDomainResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p CustomDomainResourceOperationPredicate) Matches(input CustomDomainResource) bool {

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

type DeploymentResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p DeploymentResourceOperationPredicate) Matches(input DeploymentResource) bool {

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

type GatewayCustomDomainResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GatewayCustomDomainResourceOperationPredicate) Matches(input GatewayCustomDomainResource) bool {

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

type GatewayResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GatewayResourceOperationPredicate) Matches(input GatewayResource) bool {

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

type GatewayRouteConfigResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p GatewayRouteConfigResourceOperationPredicate) Matches(input GatewayRouteConfigResource) bool {

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

type ResourceSkuOperationPredicate struct {
	Name         *string
	ResourceType *string
	Tier         *string
}

func (p ResourceSkuOperationPredicate) Matches(input ResourceSku) bool {

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	if p.Tier != nil && (input.Tier == nil && *p.Tier != *input.Tier) {
		return false
	}

	return true
}

type ServiceRegistryResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ServiceRegistryResourceOperationPredicate) Matches(input ServiceRegistryResource) bool {

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

type ServiceResourceOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ServiceResourceOperationPredicate) Matches(input ServiceResource) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
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

type StorageResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p StorageResourceOperationPredicate) Matches(input StorageResource) bool {

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
