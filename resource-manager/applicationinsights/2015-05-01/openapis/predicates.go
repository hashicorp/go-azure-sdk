package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnnotationOperationPredicate struct {
	AnnotationName    *string
	Category          *string
	EventTime         *string
	Id                *string
	Properties        *string
	RelatedAnnotation *string
}

func (p AnnotationOperationPredicate) Matches(input Annotation) bool {

	if p.AnnotationName != nil && (input.AnnotationName == nil || *p.AnnotationName != *input.AnnotationName) {
		return false
	}

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.EventTime != nil && (input.EventTime == nil || *p.EventTime != *input.EventTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Properties != nil && (input.Properties == nil || *p.Properties != *input.Properties) {
		return false
	}

	if p.RelatedAnnotation != nil && (input.RelatedAnnotation == nil || *p.RelatedAnnotation != *input.RelatedAnnotation) {
		return false
	}

	return true
}

type ApplicationInsightsComponentAPIKeyOperationPredicate struct {
	ApiKey      *string
	CreatedDate *string
	Id          *string
	Name        *string
}

func (p ApplicationInsightsComponentAPIKeyOperationPredicate) Matches(input ApplicationInsightsComponentAPIKey) bool {

	if p.ApiKey != nil && (input.ApiKey == nil || *p.ApiKey != *input.ApiKey) {
		return false
	}

	if p.CreatedDate != nil && (input.CreatedDate == nil || *p.CreatedDate != *input.CreatedDate) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}

type WorkItemConfigurationOperationPredicate struct {
	ConfigDisplayName *string
	ConfigProperties  *string
	ConnectorId       *string
	Id                *string
	IsDefault         *bool
}

func (p WorkItemConfigurationOperationPredicate) Matches(input WorkItemConfiguration) bool {

	if p.ConfigDisplayName != nil && (input.ConfigDisplayName == nil || *p.ConfigDisplayName != *input.ConfigDisplayName) {
		return false
	}

	if p.ConfigProperties != nil && (input.ConfigProperties == nil || *p.ConfigProperties != *input.ConfigProperties) {
		return false
	}

	if p.ConnectorId != nil && (input.ConnectorId == nil || *p.ConnectorId != *input.ConnectorId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefault != nil && (input.IsDefault == nil || *p.IsDefault != *input.IsDefault) {
		return false
	}

	return true
}
