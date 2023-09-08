package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceVisualizationOperationPredicate struct {
	ContainerDisplayName *string
	ContainerType        *string
	ContainerWebUrl      *string
	MediaType            *string
	ODataType            *string
	PreviewImageUrl      *string
	PreviewText          *string
	Title                *string
	Type                 *string
}

func (p ResourceVisualizationOperationPredicate) Matches(input ResourceVisualization) bool {

	if p.ContainerDisplayName != nil && (input.ContainerDisplayName == nil || *p.ContainerDisplayName != *input.ContainerDisplayName) {
		return false
	}

	if p.ContainerType != nil && (input.ContainerType == nil || *p.ContainerType != *input.ContainerType) {
		return false
	}

	if p.ContainerWebUrl != nil && (input.ContainerWebUrl == nil || *p.ContainerWebUrl != *input.ContainerWebUrl) {
		return false
	}

	if p.MediaType != nil && (input.MediaType == nil || *p.MediaType != *input.MediaType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PreviewImageUrl != nil && (input.PreviewImageUrl == nil || *p.PreviewImageUrl != *input.PreviewImageUrl) {
		return false
	}

	if p.PreviewText != nil && (input.PreviewText == nil || *p.PreviewText != *input.PreviewText) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
