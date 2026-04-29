package metricdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricDefinitionOperationPredicate struct {
	Category            *string
	DisplayDescription  *string
	Id                  *string
	IsDimensionRequired *bool
	Namespace           *string
	ResourceId          *string
}

func (p MetricDefinitionOperationPredicate) Matches(input MetricDefinition) bool {

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.DisplayDescription != nil && (input.DisplayDescription == nil || *p.DisplayDescription != *input.DisplayDescription) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDimensionRequired != nil && (input.IsDimensionRequired == nil || *p.IsDimensionRequired != *input.IsDimensionRequired) {
		return false
	}

	if p.Namespace != nil && (input.Namespace == nil || *p.Namespace != *input.Namespace) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	return true
}

type SubscriptionScopeMetricDefinitionOperationPredicate struct {
	Category            *string
	DisplayDescription  *string
	Id                  *string
	IsDimensionRequired *bool
	Namespace           *string
	ResourceId          *string
}

func (p SubscriptionScopeMetricDefinitionOperationPredicate) Matches(input SubscriptionScopeMetricDefinition) bool {

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.DisplayDescription != nil && (input.DisplayDescription == nil || *p.DisplayDescription != *input.DisplayDescription) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDimensionRequired != nil && (input.IsDimensionRequired == nil || *p.IsDimensionRequired != *input.IsDimensionRequired) {
		return false
	}

	if p.Namespace != nil && (input.Namespace == nil || *p.Namespace != *input.Namespace) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	return true
}
